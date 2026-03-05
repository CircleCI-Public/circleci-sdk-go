package fakecircle

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type resourceClass struct {
	ID            string
	ResourceClass string
	Description   string
	Tokens        []*token
}

type token struct {
	ID            string
	Nickname      string
	ResourceClass string
	Token         string
	CreatedAt     string
}

type runner struct {
	Name           string
	Hostname       string
	IP             string
	Version        string
	Status         string
	ResourceClass  string
	FirstConnected string
	LastConnected  string
	LastUsed       string
}

var (
	resourceClasses = make(map[string]*resourceClass)
	tokens          = make(map[string]*token)
	runners         = make([]*runner, 0)
)

func (s *Service) setupRunnerRoutes(r *gin.Engine) {
	r.GET("/api/v3/runner", s.listRunners)
	r.GET("/api/v3/runner/resource", s.listResourceClasses)
	r.POST("/api/v3/runner/resource", s.createResourceClass)
	r.DELETE("/api/v3/runner/resource/:id", s.deleteResourceClass)
	r.DELETE("/api/v3/runner/resource/:id/force", s.deleteResourceClassForce)
	r.GET("/api/v3/runner/token", s.listTokens)
	r.POST("/api/v3/runner/token", s.createToken)
	r.DELETE("/api/v3/runner/token/:id", s.deleteToken)
	r.GET("/api/v3/runner/tasks", s.getUnclaimedTaskCount)
	r.GET("/api/v3/runner/tasks/running", s.getRunningTaskCount)
}

func (s *Service) AddResourceClass(uuid, resourceClassIn, description string) error {
	if _, exists := resourceClasses[uuid]; exists {
		return fmt.Errorf("resourceClass with id %s already exists.", uuid)
	}
	resourceClasses[uuid] = &resourceClass{
		ID:            uuid,
		ResourceClass: resourceClassIn,
		Description:   description,
	}
	return nil
}

func (s *Service) listRunners(c *gin.Context) {
	resourceClass := c.Query("resource-class")
	namespace := c.Query("namespace")
	orgID := c.Query("org-id")

	filtered := make([]runner, 0)
	for _, r := range runners {
		if resourceClass != "" && r.ResourceClass != resourceClass {
			continue
		}
		if namespace != "" {
			// Simple namespace matching
			continue
		}
		if orgID != "" {
			// Simple org-id matching
			continue
		}
		filtered = append(filtered, *r)
	}

	c.JSON(http.StatusOK, filtered)
}

func (s *Service) listResourceClasses(c *gin.Context) {
	namespace := c.Query("namespace")
	orgID := c.Query("org-id")

	type responseItem struct {
		ID            string `json:"id"`
		ResourceClass string `json:"resource_class"`
		Description   string `json:"description"`
	}

	type response struct {
		Items []responseItem `json:"items"`
	}

	filtered := make([]responseItem, 0)
	for _, rc := range resourceClasses {
		// Simple filtering - in a real implementation this would be more sophisticated
		if namespace != "" || orgID != "" {
			filtered = append(filtered, responseItem{
				ID:            rc.ID,
				ResourceClass: rc.ResourceClass,
				Description:   rc.Description,
			})
		}
	}

	resp := response{
		Items: filtered,
	}

	c.JSON(http.StatusOK, resp)
}

func (s *Service) createResourceClass(c *gin.Context) {
	type request struct {
		OrganizationID string `json:"org_id" binding:"required"`
		ResourceClass  string `json:"resource_class" binding:"required"`
		Description    string `json:"description"`
	}

	type responseItem struct {
		ID            string `json:"id"`
		ResourceClass string `json:"resource_class"`
		Description   string `json:"description"`
	}

	type response struct {
		Items []responseItem `json:"items"`
	}

	var body request
	if err := c.BindJSON(&body); err != nil {
		msg(c, http.StatusBadRequest, "bad request")
		return
	}

	// Check for duplicates
	for _, rc := range resourceClasses {
		if rc.ResourceClass == body.ResourceClass {
			msg(c, http.StatusConflict, "resource class already exists")
			return
		}
	}

	id := uuid.New().String()
	rc := &resourceClass{
		ID:            id,
		ResourceClass: body.ResourceClass,
		Description:   body.Description,
		Tokens:        make([]*token, 0),
	}
	resourceClasses[id] = rc

	items := []responseItem{
		{
			ID:            rc.ID,
			ResourceClass: rc.ResourceClass,
			Description:   rc.Description,
		},
	}

	c.JSON(http.StatusOK, response{
		Items: items,
	})
}

func (s *Service) deleteResourceClass(c *gin.Context) {
	id := c.Param("id")

	rc, ok := resourceClasses[id]
	if !ok {
		msg(c, http.StatusNotFound, "resource class not found")
		return
	}

	// Check if there are tokens
	if len(rc.Tokens) > 0 {
		msg(c, http.StatusConflict, "resource class has tokens")
		return
	}

	delete(resourceClasses, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (s *Service) deleteResourceClassForce(c *gin.Context) {
	id := c.Param("id")

	rc, ok := resourceClasses[id]
	if !ok {
		msg(c, http.StatusNotFound, "resource class not found")
		return
	}

	// Delete all tokens
	for _, t := range rc.Tokens {
		delete(tokens, t.ID)
	}

	delete(resourceClasses, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (s *Service) listTokens(c *gin.Context) {
	resourceClass := c.Query("resource-class")

	type response struct {
		ID            string `json:"id"`
		Nickname      string `json:"nickname"`
		ResourceClass string `json:"resource_class"`
		CreatedAt     string `json:"created_at"`
	}

	filtered := make([]response, 0)
	for _, t := range tokens {
		if resourceClass == "" || t.ResourceClass == resourceClass {
			filtered = append(filtered, response{
				ID:            t.ID,
				Nickname:      t.Nickname,
				ResourceClass: t.ResourceClass,
				CreatedAt:     t.CreatedAt,
			})
		}
	}

	c.JSON(http.StatusOK, filtered)
}

func (s *Service) createToken(c *gin.Context) {
	type request struct {
		OrganizationID string `json:"org_id" binding:"required"`
		ResourceClass  string `json:"resource_class" binding:"required"`
		Nickname       string `json:"nickname" binding:"required"`
	}

	type response struct {
		ID            string `json:"id"`
		Nickname      string `json:"nickname"`
		ResourceClass string `json:"resource_class"`
		Token         string `json:"token"`
		CreatedAt     string `json:"created_at"`
	}

	var body request
	if err := c.BindJSON(&body); err != nil {
		msg(c, http.StatusBadRequest, "bad request")
		return
	}

	// Find resource class
	var rc *resourceClass
	for _, r := range resourceClasses {
		if r.ResourceClass == body.ResourceClass {
			rc = r
			break
		}
	}

	if rc == nil {
		msg(c, http.StatusNotFound, "resource class not found")
		return
	}

	id := uuid.New().String()
	tokenValue := "token_" + uuid.New().String()
	t := &token{
		ID:            id,
		Nickname:      body.Nickname,
		ResourceClass: body.ResourceClass,
		Token:         tokenValue,
		CreatedAt:     time.Now().Format(time.RFC3339),
	}
	tokens[id] = t
	rc.Tokens = append(rc.Tokens, t)

	c.JSON(http.StatusOK, response{
		ID:            t.ID,
		Nickname:      t.Nickname,
		ResourceClass: t.ResourceClass,
		Token:         t.Token,
		CreatedAt:     t.CreatedAt,
	})
}

func (s *Service) deleteToken(c *gin.Context) {
	id := c.Param("id")

	t, ok := tokens[id]
	if !ok {
		msg(c, http.StatusNotFound, "token not found")
		return
	}

	// Remove from resource class
	for _, rc := range resourceClasses {
		if rc.ResourceClass == t.ResourceClass {
			rc.Tokens = slices.DeleteFunc(rc.Tokens, func(tok *token) bool {
				return tok.ID == id
			})
			break
		}
	}

	delete(tokens, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (s *Service) getUnclaimedTaskCount(c *gin.Context) {
	type response struct {
		UnclaimedTaskCount int `json:"unclaimed_task_count"`
	}

	c.JSON(http.StatusOK, response{
		UnclaimedTaskCount: 0,
	})
}

func (s *Service) getRunningTaskCount(c *gin.Context) {
	type response struct {
		RunningRunnerTasks int `json:"running_runner_tasks"`
	}

	c.JSON(http.StatusOK, response{
		RunningRunnerTasks: 0,
	})
}
