package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"

	"ineedApp/app/models"
)

type BusinessesResource struct {
	buffalo.Resource
}

func (v BusinessesResource) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	businesses := models.Businesses{}

	q := tx.PaginateFromParams(c.Params())
	if err := q.All(&businesses); err != nil {
		return err
	}

	defaults := models.Businesses{
		{Name: "Beauty Nails", Description: "Get your nails done quickly! We're a professional team of fashionist.", Phone: "3045421621", Img: "nail", Stars: 0, Category: "Style", ServiceTime: "Mon to Fri from 9:00am to 6:00pm"},
		{Name: "Plumber 24/7", Description: "Got a leak?, we're a 24 hourse plumber team. Synk, Dishwasher, Toilet and all kind", Phone: "3045421621", Img: "plumber", Stars: 4, Category: "Home Needs"},
		{Name: "FIXAQUA - Barranquilla", Description: "More than 25 years in the bussiness fixing plumb issues", Phone: "3045421621", Img: "plumber", Stars: 5, Category: "Home Needs"},
		{Name: "Pool's Clean", Description: "Cleaning pool's since 1966, contact our service now!", Phone: "3045421621", Img: "pool", Stars: 5, Category: "Home Needs"},
		{Name: "Vikings Look Barbershop", Description: "The best and moderns haircuts for men.", Phone: "3045421621", Img: "barber_shop", Stars: 5, Category: "Style"},
		{Name: "Makeupme Over", Description: "MakeUp experts, for all kind of events, hire us.", Phone: "3045421621", Img: "makeup", Stars: 5, Category: "Style"},
		{Name: "MB Eye Brows", Description: "Microblading for thickert eye brows, lashes all kind, shape and color", Phone: "3045421621", Img: "makeup", Stars: 5, Category: "Style"},
		{Name: "Barberin", Description: "I'm a experienced barber with more than 10 years of experience.", Phone: "3045421621", Img: "barber_shop", Stars: 0, Category: "Style"},
	}

	businesses = append(businesses, defaults...)

	c.Set("businesses", businesses)

	return c.Render(http.StatusOK, r.HTML("/business/index.plush.html"))
}

func (v BusinessesResource) ListBussines(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	businesses := models.Businesses{}

	q := tx.PaginateFromParams(c.Params())
	if err := q.All(&businesses); err != nil {
		return err
	}

	defaults := models.Businesses{
		{Name: "Beauty Nails", Description: "Get your nails done quickly! We're a professional team of fashionist.", Phone: "3045421621", Img: "nail", Stars: 0, Category: "Style", ServiceTime: "Mon to Fri from 9:00am to 6:00pm"},
		{Name: "Plumber 24/7", Description: "Got a leak?, we're a 24 hourse plumber team. Synk, Dishwasher, Toilet and all kind", Phone: "3045421621", Img: "plumber", Stars: 4, Category: "Home Needs"},
		{Name: "FIXAQUA - Barranquilla", Description: "More than 25 years in the bussiness fixing plumb issues", Phone: "3045421621", Img: "plumber", Stars: 5, Category: "Home Needs"},
		{Name: "Pool's Clean", Description: "Cleaning pool's since 1966, contact our service now!", Phone: "3045421621", Img: "pool", Stars: 5, Category: "Home Needs"},
		{Name: "Vikings Look Barbershop", Description: "The best and moderns haircuts for men.", Phone: "3045421621", Img: "barber_shop", Stars: 5, Category: "Style"},
		{Name: "Makeupme Over", Description: "MakeUp experts, for all kind of events, hire us.", Phone: "3045421621", Img: "makeup", Stars: 5, Category: "Style"},
		{Name: "MB Eye Brows", Description: "Microblading for thickert eye brows, lashes all kind, shape and color", Phone: "3045421621", Img: "makeup", Stars: 5, Category: "Style"},
		{Name: "Barberin", Description: "I'm a experienced barber with more than 10 years of experience.", Phone: "3045421621", Img: "barber_shop", Stars: 0, Category: "Style"},
	}

	businesses = append(businesses, defaults...)

	c.Set("businesses", businesses)
	return c.Render(http.StatusOK, r.JSON(businesses))
}

func (v BusinessesResource) Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	business := &models.Business{}
	if err := tx.Find(business, c.Param("business_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("business", business)

	return c.Render(http.StatusOK, r.HTML("/business/show.plush.html"))
}

func (v BusinessesResource) GetServiceForBusiness(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	services := &models.Services{}

	if err := tx.Where("business_id = ?", c.Param("business_id")).All(services); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("services", services)
	return c.Render(http.StatusOK, r.HTML("/business/show.plush.html"))
}

func (v BusinessesResource) New(c buffalo.Context) error {
	c.Set("business", &models.Business{})
	c.Set("categories", map[string]string{"Style": "Style", "Health": "Health", "Home Needs": "Home Needs"})

	return c.Render(http.StatusOK, r.HTML("/business/new.plush.html"))
}

func (v BusinessesResource) Create(c buffalo.Context) error {
	business := models.Business{}

	if err := c.Bind(&business); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	business.Address = "Default address"

	verrs, err := tx.ValidateAndCreate(&business)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("categories", map[string]string{"Style": "Style", "Health": "Health"})
		c.Set("errors", verrs)
		c.Set("business", business)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/business/new.plush.html"))
	}

	businesses := models.Businesses{}

	q := tx.PaginateFromParams(c.Params())
	if err := q.All(&businesses); err != nil {
		return err
	}

	c.Set("businesses", businesses)

	c.Flash().Add("success", "business.created.success")
	return c.Redirect(http.StatusSeeOther, "/business/list")
	// return c.Render(http.StatusOK, r.HTML("/business/index.plush.html"))
}

func (v BusinessesResource) Edit(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	business := &models.Business{}

	if err := tx.Find(business, c.Param("business_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("business", business)
	return c.Render(http.StatusOK, r.HTML("/business/edit.plush.html"))
}

func (v BusinessesResource) Update(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	business := &models.Business{}

	if err := tx.Find(business, c.Param("business_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := c.Bind(business); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(business)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("business", business)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/business/edit.plush.html"))
	}

	c.Flash().Add("success", "business.updated.success")
	return c.Redirect(http.StatusSeeOther, "/business/list")
}

// Destroy deletes a Business from the DB. This function is mapped
// to the path DELETE /businesses/{business_id}
func (v BusinessesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Business
	business := &models.Business{}

	// To find the Business the parameter business_id is used.
	if err := tx.Find(business, c.Param("business_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(business); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "business.destroyed.success")

	// Redirect to the index page
	return c.Redirect(http.StatusSeeOther, "businessesPath()")
}
