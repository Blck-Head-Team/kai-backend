package newsletters

import (
	"kai-backend/business/newsletters"
	"kai-backend/business/paginations"
	"kai-backend/controllers"
	requests "kai-backend/controllers/newsletters/request"
	responses "kai-backend/controllers/newsletters/response"
	"kai-backend/exceptions"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	Usecase newsletters.Usecase
}

type Header struct {
	Cookie string `json:"cookie"`
}

func NewHandler(n newsletters.Usecase) *Controllers {
	return &Controllers{
		Usecase: n,
	}
}

func (n *Controllers) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	paginationDomain := paginations.Domain{
		Page:  1,
		Limit: 0,
	}

	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	sort := c.QueryParam("sort")

	var intPage, intLimit int
	var err error
	if page != "" {
		intPage, err = strconv.Atoi(page)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}
		paginationDomain.Page = intPage
	}
	if limit != "" {
		intLimit, err = strconv.Atoi(limit)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}
		paginationDomain.Limit = intLimit
	}

	paginationDomain.Sort = sort

	newsletters, err := n.Usecase.GetAll(ctx, paginationDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := make([]responses.NewsResponse, len(newsletters))
	for i, news := range newsletters {
		response[i] = responses.FromDomain(news)
	}
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (n *Controllers) CountAll(c echo.Context) error {
	ctx := c.Request().Context()

	count, err := n.Usecase.CountAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, http.StatusOK, count)
}

func (n *Controllers) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("Id")
	class, err := n.Usecase.GetById(ctx, id)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrNewsNotFound)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := responses.FromDomain(class)
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (n *Controllers) Create(c echo.Context) error {
	ctx := c.Request().Context()

	createdNews := requests.CreateNews{}
	c.Bind(&createdNews)

	newsDomain := newsletters.Domain{
		Title:       createdNews.Title,
		Description: createdNews.Description,
		Content:     createdNews.Content,
	}

	news, err := n.Usecase.Create(ctx, newsDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := responses.FromDomain(news)
	return controllers.SuccessResponse(c, http.StatusCreated, response)
}

func (n *Controllers) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("Id")
	updatedNews := requests.CreateNews{}
	c.Bind(&updatedNews)

	updatedNewsDomain := newsletters.Domain{
		Title:       updatedNews.Title,
		Description: updatedNews.Description,
		Content:     updatedNews.Content,
	}

	news, err := n.Usecase.Update(ctx, id, updatedNewsDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := responses.FromDomain(news)
	return controllers.SuccessResponse(c, http.StatusOK, response)
}

func (n *Controllers) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("Id")
	err := n.Usecase.Delete(ctx, id)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return controllers.ErrorResponse(c, http.StatusNotFound, exceptions.ErrClassNotFound)
		}
		if err == exceptions.ErrEmptyInput {
			return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrEmptyInput)
		}
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}
