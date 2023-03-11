package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/pop/v6"

	"ineedApp/app/models"
)

// ClientsResource is the resource for the Client model
type ClientsResource struct {
	buffalo.Resource
}

// List gets all Clients. This function is mapped to the path
// GET /clients
func (v ClientsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	clients := &models.Clients{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Clients from the DB
	if err := q.All(clients); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	c.Set("clients", clients)

	return c.Render(http.StatusOK, r.HTML("/clients/index.plush.html"))
}

// Show gets the data for one Client. This function is mapped to
// the path GET /clients/{client_id}
func (v ClientsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Client
	client := &models.Client{}

	// To find the Client the parameter client_id is used.
	if err := tx.Find(client, c.Param("client_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("client", client)

	return c.Render(http.StatusOK, r.HTML("/clients/show.plush.html"))
}

// New renders the form for creating a new Client.
// This function is mapped to the path GET /clients/new
func (v ClientsResource) New(c buffalo.Context) error {
	c.Set("client", &models.Client{})

	return c.Render(http.StatusOK, r.HTML("/clients/new.plush.html"))
}

// Create adds a Client to the DB. This function is mapped to the
// path POST /clients
func (v ClientsResource) Create(c buffalo.Context) error {
	// Allocate an empty Client
	client := &models.Client{}

	// Bind client to the html form elements
	if err := c.Bind(client); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(client)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		c.Set("client", client)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/clients/new.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "client.created.success")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "clientPath()", render.Data{"client_id": client.ID})
}

// Edit renders a edit form for a Client. This function is
// mapped to the path GET /clients/{client_id}/edit
func (v ClientsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Client
	client := &models.Client{}

	if err := tx.Find(client, c.Param("client_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("client", client)

	return c.Render(http.StatusOK, r.HTML("/clients/edit.plush.html"))
}

// Update changes a Client in the DB. This function is mapped to
// the path PUT /clients/{client_id}
func (v ClientsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Client
	client := &models.Client{}

	if err := tx.Find(client, c.Param("client_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Client to the html form elements
	if err := c.Bind(client); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(client)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		c.Set("client", client)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/clients/edit.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "client.updated.success")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "clientPath()", render.Data{"client_id": client.ID})
}

// Destroy deletes a Client from the DB. This function is mapped
// to the path DELETE /clients/{client_id}
func (v ClientsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Client
	client := &models.Client{}

	// To find the Client the parameter client_id is used.
	if err := tx.Find(client, c.Param("client_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(client); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "client.destroyed.success")

	// Redirect to the index page
	return c.Redirect(http.StatusSeeOther, "clientsPath()")
}

func (v ClientsResource) GetServiceForClients(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	services := &models.Services{}

	if err := tx.Where("client_id = ?", c.Param("client_id")).All(services); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("services", services)
	return c.Render(http.StatusOK, r.HTML("/business/show.plush.html"))
}
