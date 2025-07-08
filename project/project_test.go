package project_test

import (
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/fakecircle"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/integrationtest"
	"github.com/CircleCI-Public/circleci-sdk-go/project"
)

const testTok = "8f23dc1b-b7fd-4bed-9a2c-ec699b1ba810"

func TestProjectService_Get(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	ps := project.NewProjectService(c)

	org, err := fc.AddOrg(fakecircle.NewOrg{
		Type: fakecircle.TypeCircleCI,
		Name: "test org",
	})
	assert.Assert(t, err)
	prj, err := fc.AddProject(fakecircle.NewProject{
		OrgID: org.ID,
		Name:  "test project",
	})
	assert.Assert(t, err)

	t.Run("get", func(t *testing.T) {
		gotProj, err := ps.Get(prj.Slug)
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(gotProj, &project.Project{
			Id:               prj.ID.String(),
			Name:             "test project",
			Slug:             prj.Slug,
			OrganizationName: "test org",
			OrganizationSlug: org.Slug,
			OrganizationId:   org.ID.String(),
			VcsInfo: common.VcsInfo{
				VcsUrl:        "git://github.com/dummy-value",
				Provider:      fakecircle.TypeCircleCI,
				DefaultBranch: "main",
			},
		}))
	})
}

func TestProjectService_Get_Integration(t *testing.T) {
	c := integrationtest.Client(t)
	projectService := project.NewProjectService(c)

	slug := "circleci/8e4z1Akd74woxagxnvLT5q/V29Cenkg8EaiSZARmWm8Lz"
	p, err := projectService.Get(slug)
	assert.Assert(t, err)

	t.Log(p)
	assert.Check(t, cmp.Equal(p.Slug, slug))
}

func TestProjectService_Create(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	ps := project.NewProjectService(c)

	org, err := fc.AddOrg(fakecircle.NewOrg{
		Type: fakecircle.TypeCircleCI,
		Name: "test org",
	})
	assert.Assert(t, err)

	var p *project.Project
	t.Run("create", func(t *testing.T) {
		var err error
		p, err = ps.Create("test project name", org.ID.String())
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(p, &project.Project{
			Id:               "ignored",
			Name:             "test project name",
			Slug:             "ignored",
			OrganizationName: "test org",
			OrganizationSlug: org.Slug,
			OrganizationId:   org.ID.String(),
			VcsInfo: common.VcsInfo{
				VcsUrl:        "git://github.com/dummy-value",
				Provider:      fakecircle.TypeCircleCI,
				DefaultBranch: "main",
			},
		}, cmpopts.IgnoreFields(project.Project{}, "Id", "Slug")))
	})

	t.Run("get", func(t *testing.T) {
		p, err := fc.Project(uuid.MustParse(p.Id))
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(p, fakecircle.Project{
			ID:   p.ID,
			Name: "test project name",
			Slug: p.Slug,
			Org: fakecircle.Org{
				ID:   p.Org.ID,
				Type: fakecircle.TypeCircleCI,
				Name: "test org",
				Slug: p.Org.Slug,
			},
		}))
	})
}

func TestProjectService_Delete(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	ps := project.NewProjectService(c)

	org, err := fc.AddOrg(fakecircle.NewOrg{
		Type: fakecircle.TypeCircleCI,
		Name: "test org",
	})
	assert.Assert(t, err)
	prj, err := fc.AddProject(fakecircle.NewProject{
		OrgID: org.ID,
		Name:  "test project",
	})
	assert.Assert(t, err)

	t.Run("delete", func(t *testing.T) {
		err := ps.Delete(prj.Slug)
		assert.Assert(t, err)
	})

	t.Run("get", func(t *testing.T) {
		p, err := fc.Project(prj.ID)
		assert.Check(t, cmp.ErrorContains(err, "not found"))
		assert.Check(t, cmp.Equal(p.ID, uuid.Nil))
	})
}

func TestFullProject_Integration(t *testing.T) {
	c := integrationtest.Client(t)
	projectService := project.NewProjectService(c)

	name := "test-api-client-repo"
	organizationID := "3ddcf1d1-7f5f-4139-8cef-71ad0921a968"

	p, err := projectService.Create(name, organizationID)
	assert.Assert(t, err)
	t.Log(p)

	projectSettings, err := projectService.GetSettings(fakecircle.TypeCircleCI, p.OrganizationId, p.Id)
	assert.Assert(t, err)
	t.Log(projectSettings)
	newSettings := project.ProjectSettings{
		Advanced: project.AdvanceSettings{
			AutocancelBuilds: common.Bool(true),
			DisableSSH:       common.Bool(false),
		},
	}
	projectSettings, err = projectService.UpdateSettings(newSettings, fakecircle.TypeCircleCI, p.OrganizationId, p.Id)
	assert.Assert(t, err)
	t.Log(projectSettings)

	err = projectService.Delete(p.Slug)
	assert.Assert(t, err)
}

func TestClassicProject(t *testing.T) {
	t.Skip("This test is for manual usage only")

	c := integrationtest.Client(t)
	projectService := project.NewProjectService(c)

	name := "test-api-client-repo"
	organizationID := "30361eb7-0812-447f-bca5-a299d96576c7"
	_, err := projectService.Create(name, organizationID)
	assert.Assert(t, err)
}

func TestDeleteProject(t *testing.T) {
	t.Skip("This test is for manual usage only")

	c := integrationtest.Client(t)
	projectService := project.NewProjectService(c)

	err := projectService.Delete("github/marboledacci/test-api-client-repo")
	assert.Assert(t, err)
}
