package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/pop/v6"

	"ineedApp/app/models"
)

// ServicesResource is the resource for the Service model
type ServicesResource struct {
	buffalo.Resource
}

// List gets all Services. This function is mapped to the path
// GET /services
func (v ServicesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	services := &models.Services{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Services from the DB
	if err := q.All(services); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	c.Set("services", services)

	return c.Render(http.StatusOK, r.HTML("/services/index.plush.html"))
}

// Show gets the data for one Service. This function is mapped to
// the path GET /services/{service_id}
func (v ServicesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Service
	service := &models.Service{}

	// To find the Service the parameter service_id is used.
	if err := tx.Find(service, c.Param("service_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("service", service)

	return c.Render(http.StatusOK, r.HTML("/services/show.plush.html"))
}

// New renders the form for creating a new Service.
// This function is mapped to the path GET /services/new
func (v ServicesResource) New(c buffalo.Context) error {
	c.Set("service", &models.Service{})

	return c.Render(http.StatusOK, r.HTML("/services/new.plush.html"))
}

// Create adds a Service to the DB. This function is mapped to the
// path POST /services
func (v ServicesResource) Create(c buffalo.Context) error {
	// Allocate an empty Service
	service := &models.Service{}

	// Bind service to the html form elements
	if err := c.Bind(service); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(service)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		c.Set("service", service)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/services/new.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "service.created.success")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "servicePath()", render.Data{"service_id": service.ID})
}

// Edit renders a edit form for a Service. This function is
// mapped to the path GET /services/{service_id}/edit
func (v ServicesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Service
	service := &models.Service{}

	if err := tx.Find(service, c.Param("service_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("service", service)

	return c.Render(http.StatusOK, r.HTML("/services/edit.plush.html"))
}

// Update changes a Service in the DB. This function is mapped to
// the path PUT /services/{service_id}
func (v ServicesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Service
	service := &models.Service{}

	if err := tx.Find(service, c.Param("service_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Service to the html form elements
	if err := c.Bind(service); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(service)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		c.Set("service", service)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/services/edit.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "service.updated.success")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "servicePath()", render.Data{"service_id": service.ID})
}

// Destroy deletes a Service from the DB. This function is mapped
// to the path DELETE /services/{service_id}
func (v ServicesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Service
	service := &models.Service{}

	// To find the Service the parameter service_id is used.
	if err := tx.Find(service, c.Param("service_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(service); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "service.destroyed.success")

	// Redirect to the index page
	return c.Redirect(http.StatusSeeOther, "servicesPath()")
}
