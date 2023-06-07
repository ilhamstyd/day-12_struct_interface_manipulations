package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	ProjectName string
	StartDate   string
	EndDate     string
	Duration    string
	Description string
	Author      string
}

var dataProject = []Project{
	{
		ProjectName: "gak punya duit rek",
		StartDate:   "07/06/2023",
		EndDate:     "08/06/2023",
		Duration:    "1 hari",
		Description: "punya duit pusing gak punya duit lebih pusing",
		Author:      "wa doyok",
	},
	{
		ProjectName: "gak punya duit rek",
		StartDate:   "07/06/2023",
		EndDate:     "08/06/2023",
		Duration:    "1 hari",
		Description: "punya duit pusing gak punya duit lebih pusing",
		Author:      "wa bewok",
	},
	{
		ProjectName: "gak punya duit rek",
		StartDate:   "07/06/2023",
		EndDate:     "09/06/2023",
		Duration:    "3 hari",
		Description: "punya duit pusing gak punya duit lebih pusing",
		Author:      "wa kumis",
	},
}

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/hello", helloword)
	e.GET("/", home)
	e.GET("/addProject", addProject)
	e.GET("/projeect-detail/:id", projectDetail)
	e.GET("/contactMe", contactMe)

	e.POST("/edit-project/:id", editProject)
	e.POST("/delete-project/:id", deleteProject)
	e.POST("/addFormProject", addFormProject)

	e.Logger.Fatal(e.Start("localhost:8000"))
}

func helloword(c echo.Context) error {
	return c.String(http.StatusOK, "helloworld")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"projects": dataProject,
	}

	return tmpl.Execute(c.Response(), projects)
}

func addProject(c echo.Context) error {
	var template, err = template.ParseFiles("views/addProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return template.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for index, item := range dataProject {
		if id == index {
			ProjectDetail = Project{
				ProjectName: item.ProjectName,
				StartDate:   item.StartDate,
				EndDate:     item.EndDate,
				Duration:    item.Duration,
				Description: item.Description,
				Author:      item.Author,
			}
		}
	}

	item := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var template, err = template.ParseFiles("views/add-project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return template.Execute(c.Response(), item)
}

func contactMe(c echo.Context) error {
	var template, err = template.ParseFiles("views/contact-me.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return template.Execute(c.Response(), nil)
}

func addFormProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("desc")

	println("Project name : " + projectName)
	println("start date : " + startDate)
	println("end date : " + endDate)
	println("description : " + description)

	var newProject = Project{
		ProjectName: projectName,
		StartDate:   startDate,
		EndDate:     endDate,
		Duration:    "1 hari",
		Description: description,
		Author:      "ASTA",
	}

	dataProject = append(dataProject, newProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(delete echo.Context) error {
	i, _ := strconv.Atoi(delete.Param("id"))

	fmt.Println("index : ", i)

	dataProject = append(dataProject[:i], dataProject[i+1:]...)

	return delete.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(edit echo.Context) error {
	id, _ := strconv.Atoi(edit.Param("id"))
	fmt.Println("index : ", id)

	dataProject = append(dataProject[:id], dataProject[id+1:]...)
	return edit.Redirect(http.StatusMovedPermanently, "/addProject")
}
