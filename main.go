package main

import (
	"b47s1/connection"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)
type Project struct {
	id          int
	projectName string
	startDate   string
	endDate     string
	duration    string
	description string
	nodeJs      bool
	reactJs     bool
	nextJs      bool
	typescript  bool
}
type user struct {
	ID int
	Name string
	Email string
	password string
}

// var dataProject = []Project{
// 	{
// 		ProjectName: "case 1",
// 		StartDate:   "2023-07-01",
// 		EndDate:     "2023-08-01",
// 		Duration:    "1 Bulan",
// 		Description: "contoh 1",
// 		nodeJs: true,
// 		reactJs: true,
// 		nextJs: true,
// 		typescript: true,
// 	},
// 	{
// 		ProjectName: "case 2",
// 		StartDate:   "2023-07-01",
// 		EndDate:     "2023-08-01",
// 		Duration:    "1 Bulan",
// 		Description: "contoh 2",
// 		nodeJs: true,
// 		reactJs: false,
// 		nextJs: true,
// 		typescript: false,
// 	},
// 	{
// 		ProjectName: "case 3",
// 		StartDate:   "2023-07-01",
// 		EndDate:     "2023-08-01",
// 		Duration:    "1 Bulan",
// 		Description: "contoh 3",
// 		nodeJs: true,
// 		reactJs: true,
// 		nextJs: true,
// 		typescript: false,
// 	},
// 	{
// 		ProjectName: "case 4",
// 		StartDate:   "2023-07-01",
// 		EndDate:     "2023-08-01",
// 		Duration:    "1 Bulan",
// 		Description: "contoh 4",
// 		nodeJs: false,
// 		reactJs: false,
// 		nextJs: true,
// 		typescript: false,
// 	},
// 	{
// 		ProjectName: "case 4",
// 		StartDate:   "2023-07-01",
// 		EndDate:     "2023-08-01",
// 		Duration:    "1 Bulan",
// 		Description: "contoh 4",
// 		nodeJs: true,
// 		reactJs: false,
// 		nextJs: true,
// 		typescript: false,
// 	},
// }
func main() {
	connection.DatabaseConnect()
	e := echo.New()

	e.Static("/public", "public")
//routing

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/testimonials", testimonials)
	e.GET("/project/:id", projectDetail)
	e.GET("/addProject", addProject)
	e.GET("/editProject/:id", editProject)
	
	
	//post


e.POST("/editproject/:id", editProject)
e.POST("/project-delete/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))

}
func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, projectName, startDate, endDate, duration, description, nodeJs, reactJs, nextJs, typescript FROM tb_blog ORDER BY id ASC")

	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.id, &each.projectName, &each.startDate, &each.endDate, &each.duration, &each.description, &each.nodeJs, &each.reactJs, &each.nextJs, &each.typescript)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		result = append(result, each)
	}

	fmt.Println(result)

	projects := map[string]interface{}{
		"Projects": result,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messsage": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, projectName, startDate, endDate, duration, description, nodeJs, reactJs, nextJs, typescript, image FROM tb_blog WHERE id=$1", id).Scan(
		&ProjectDetail.id, &ProjectDetail.projectName, &ProjectDetail.startDate, &ProjectDetail.endDate, &ProjectDetail.duration, &ProjectDetail.description, &ProjectDetail.nodeJs, &ProjectDetail.reactJs, &ProjectDetail.nextJs, &ProjectDetail.typescript)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/project-detail.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, projectName, startDate, endDate, duration, description, nodeJs, reactJs, nextJs, typescript, image FROM tb_blog WHERE id=$1", id).Scan(&ProjectDetail.id, &ProjectDetail.projectName, &ProjectDetail.startDate, &ProjectDetail.endDate, &ProjectDetail.duration, &ProjectDetail.description, &ProjectDetail.nodeJs, &ProjectDetail.reactJs, &ProjectDetail.nextJs, &ProjectDetail.typescript)

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, errTmplt = template.ParseFiles("views/edit-project.html")

	if errTmplt != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := hitungDurasi(startDate, endDate)
	description := c.FormValue("description")
	nodeJs := (c.FormValue("nodeJs") == "nodeJs")
	reactJs := (c.FormValue("reactJs") == "reactJs")
	nextJs := (c.FormValue("nextJs") == "nextJs")
	typescript := (c.FormValue("typescript") == "typescript")
	image := c.FormValue("input-image")

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (projectName, startDate, endDate, duration, description, nodeJs, reactJs, nextJs, typescript, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", projectName, startDate, endDate, duration, description, nodeJs, reactJs,nextJs , typescript, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index :", id)

	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := hitungDurasi(startDate, endDate)
	description := c.FormValue("description")
	nodeJs := (c.FormValue("nodeJs") == "nodeJs")
	reactJs := (c.FormValue("reactJs") == "reactJs")
	nextJs := (c.FormValue("nextJs") == "nextJs")
	typescript := (c.FormValue("typescript") == "typescript")
	image := c.FormValue("input-image")

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_blog SET projectName=$1, startDate=$2, endDate=$3, duration=$4, description=$5, nodeJs=$6, reactJs=$7, nextJs=$8, typescript=$9, image=$10 WHERE id=$11", projectName, startDate, endDate, duration, description, nodeJs, reactJs, nextJs, typescript, image, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func hitungDurasi(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " Tahun"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " Tahun"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " Bulan"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " Bulan"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " Minggu"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " Minggu"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " Hari"
				} else {
					duration = strconv.Itoa(durationDays) + " Hari"
				}
			}
		}
	}

	return duration
}