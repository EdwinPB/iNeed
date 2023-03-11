package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"

	"ineedApp/app/models"
)

// BusinessesResource is the resource for the Business model
type BusinessesResource struct {
	buffalo.Resource
}

// List gets all Businesses. This function is mapped to the path
// GET /businesses
func (v BusinessesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	// tx, ok := c.Value("tx").(*pop.Connection)
	// if !ok {
	// 	return fmt.Errorf("no transaction found")
	// }

	// businesses := &models.Businesses{}

	// // Paginate results. Params "page" and "per_page" control pagination.
	// // Default values are "page=1" and "per_page=20".
	// q := tx.PaginateFromParams(c.Params())

	// // Retrieve all Businesses from the DB
	// if err := q.All(businesses); err != nil {
	// 	return err
	// }

	// // Add the paginator to the context so it can be used in the template.
	// c.Set("pagination", q.Paginator)
	// c.Set("businesses", businesses)

	return c.Render(http.StatusOK, r.HTML("/business/index.plush.html"))
}

func (v BusinessesResource) ListBussines(c buffalo.Context) error {
	// Get the DB connection from the context
	// tx, ok := c.Value("tx").(*pop.Connection)
	// if !ok {
	// 	return fmt.Errorf("no transaction found")
	// }

	businesses := &models.Businesses{
		{
			ID:          uuid.Must(uuid.NewV4()),
			Name:        "Barber Shop El Calvo",
			Description: "Calle 40 # 56 - 70",
			Phone:       "606-0000-000",
		},
		{
			ID:          uuid.Must(uuid.NewV4()),
			Name:        "Barber Shop El Calvo",
			Description: "Calle 40 # 56 - 70",
			Phone:       "606-0000-000",
		},
		{
			ID:          uuid.Must(uuid.NewV4()),
			Name:        "Barber Shop El Calvo",
			Description: "Calle 40 # 56 - 70",
			Phone:       "606-0000-000",
		},
	}

	// q := tx.PaginateFromParams(c.Params())

	// if err := q.All(businesses); err != nil {
	// 	return err
	// }

	c.Set("businesses", businesses)
	return c.Render(http.StatusOK, r.JSON(businesses))
}

// Show gets the data for one Business. This function is mapped to
// the path GET /businesses/{business_id}
func (v BusinessesResource) Show(c buffalo.Context) error {
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

	c.Set("business", business)

	return c.Render(http.StatusOK, r.HTML("/businesses/show.plush.html"))
}

// New renders the form for creating a new Business.
// This function is mapped to the path GET /businesses/new
func (v BusinessesResource) New(c buffalo.Context) error {
	c.Set("business", &models.Business{})

	return c.Render(http.StatusOK, r.HTML("/businesses/new.plush.html"))
}

// Create adds a Business to the DB. This function is mapped to the
// path POST /businesses
func (v BusinessesResource) Create(c buffalo.Context) error {
	// Allocate an empty Business
	business := &models.Business{}

	// Bind business to the html form elements
	if err := c.Bind(business); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(business)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		c.Set("business", business)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/businesses/new.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "business.created.success")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "businessPath()", render.Data{"business_id": business.ID})
}

// Edit renders a edit form for a Business. This function is
// mapped to the path GET /businesses/{business_id}/edit
func (v BusinessesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Business
	business := &models.Business{}

	if err := tx.Find(business, c.Param("business_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("business", business)

	return c.Render(http.StatusOK, r.HTML("/businesses/edit.plush.html"))
}

// Update changes a Business in the DB. This function is mapped to
// the path PUT /businesses/{business_id}
func (v BusinessesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Business
	business := &models.Business{}

	if err := tx.Find(business, c.Param("business_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Business to the html form elements
	if err := c.Bind(business); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(business)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		c.Set("business", business)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/businesses/edit.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "business.updated.success")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "businessPath()", render.Data{"business_id": business.ID})
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
