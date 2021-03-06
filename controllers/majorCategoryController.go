package controllers

import(
	"fmt"
	"strconv"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
	"github.com/myrachanto/astore/service"
)
 
var (
	MCategoryController mcategoryController = mcategoryController{}
)
type mcategoryController struct{ }
/////////controllers/////////////////
func (controller mcategoryController) Create(c echo.Context) error {
	majorcategory := &model.Majorcategory{}
	
	if err := c.Bind(majorcategory); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	createdmajorcategory, err1 := service.Mcategoryservice.Create(majorcategory)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, createdmajorcategory)
}
func (controller mcategoryController) GetAll(c echo.Context) error {
	majorcategorys := []model.Majorcategory{}
	majorcategorys, err3 := service.Mcategoryservice.GetAll(majorcategorys)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, majorcategorys)
} 
func (controller mcategoryController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println(id)
	majorcategory, problem := service.Mcategoryservice.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, majorcategory)	
}

func (controller mcategoryController) Update(c echo.Context) error {
		
	majorcategory :=  &model.Majorcategory{}
	if err := c.Bind(majorcategory); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	updatedmajorcategory, problem := service.Mcategoryservice.Update(id, majorcategory)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, updatedmajorcategory)
}

func (controller mcategoryController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	success, failure := service.Mcategoryservice.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}
