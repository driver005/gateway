package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Note struct {
	r Registry
}

func NewNote(r Registry) *Note {
	m := Note{r: r}
	return &m
}

func (m *Note) SetRoutes(router fiber.Router) {
	route := router.Group("/notes")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
}

// @oas:path [get] /admin/notes/{id}
// operationId: "GetNotesNote"
// summary: "Get a Note"
// description: "Retrieve a note's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the note.
//
// x-codegen:
//
//	method: retrieve
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.notes.retrieve(noteId)
//     .then(({ note }) => {
//     console.log(note.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminNote } from "medusa-react"
//
//     type Props = {
//     noteId: string
//     }
//
//     const Note = ({ noteId }: Props) => {
//     const { note, isLoading } = useAdminNote(noteId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {note && <span>{note.resource_type}</span>}
//     </div>
//     )
//     }
//
//     export default Note
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/notes/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Notes
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminNotesRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *Note) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.NoteService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /admin/notes
// operationId: "GetNotes"
// summary: "List Notes"
// x-authenticated: true
// description: "Retrieve a list of notes. The notes can be filtered by fields such as `resource_id`. The notes can also be paginated."
// parameters:
//   - (query) limit=50 {number} Limit the number of notes returned.
//   - (query) offset=0 {number} The number of notes to skip when retrieving the notes.
//   - (query) resource_id {string} Filter by resource ID
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetNotesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.notes.list()
//     .then(({ notes, limit, offset, count }) => {
//     console.log(notes.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminNotes } from "medusa-react"
//
//     const Notes = () => {
//     const { notes, isLoading } = useAdminNotes()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {notes && !notes.length && <span>No Notes</span>}
//     {notes && notes.length > 0 && (
//     <ul>
//     {notes.map((note) => (
//     <li key={note.id}>{note.resource_type}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Notes
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/notes' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Notes
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminNotesListRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *Note) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableNote](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.NoteService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

// @oas:path [post] /admin/notes
// operationId: "PostNotes"
// summary: "Create a Note"
// description: "Create a Note which can be associated with any resource."
// x-authenticated: true
// requestBody:
// content:
//
//	application/json:
//	  schema:
//	    $ref: "#/components/schemas/AdminPostNotesReq"
//
// x-codegen:
//
//	method: create
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.notes.create({
//     resource_id,
//     resource_type: "order",
//     value: "We delivered this order"
//     })
//     .then(({ note }) => {
//     console.log(note.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateNote } from "medusa-react"
//
//     const CreateNote = () => {
//     const createNote = useAdminCreateNote()
//     // ...
//
//     const handleCreate = () => {
//     createNote.mutate({
//     resource_id: "order_123",
//     resource_type: "order",
//     value: "We delivered this order"
//     }, {
//     onSuccess: ({ note }) => {
//     console.log(note.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateNote
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/notes' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "resource_id": "{resource_id}",
//     "resource_type": "order",
//     "value": "We delivered this order"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Notes
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminNotesRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *Note) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateNoteInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.NoteService().SetContext(context.Context()).Create(model, map[string]interface{}{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/notes/{id}
// operationId: "PostNotesNote"
// summary: "Update a Note"
// x-authenticated: true
// description: "Update a Note's details."
// parameters:
//   - (path) id=* {string} The ID of the Note
//
// requestBody:
// content:
//
//	application/json:
//	  schema:
//	    $ref: "#/components/schemas/AdminPostNotesNoteReq"
//
// x-codegen:
//
//	method: update
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.notes.update(noteId, {
//     value: "We delivered this order"
//     })
//     .then(({ note }) => {
//     console.log(note.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateNote } from "medusa-react"
//
//     type Props = {
//     noteId: string
//     }
//
//     const Note = ({ noteId }: Props) => {
//     const updateNote = useAdminUpdateNote(noteId)
//     // ...
//
//     const handleUpdate = (
//     value: string
//     ) => {
//     updateNote.mutate({
//     value
//     }, {
//     onSuccess: ({ note }) => {
//     console.log(note.value)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Note
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/notes/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "value": "We delivered this order"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Notes
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminNotesRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *Note) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateNoteInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.NoteService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/notes/{id}
// operationId: "DeleteNotesNote"
// summary: "Delete a Note"
// description: "Delete a Note."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Note to delete.
//
// x-codegen:
//
//	method: delete
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.notes.delete(noteId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteNote } from "medusa-react"
//
//     type Props = {
//     noteId: string
//     }
//
//     const Note = ({ noteId }: Props) => {
//     const deleteNote = useAdminDeleteNote(noteId)
//     // ...
//
//     const handleDelete = () => {
//     deleteNote.mutate()
//     }
//
//     // ...
//     }
//
//     export default Note
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/notes/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Notes
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminNotesDeleteRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *Note) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.NoteService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "note",
		"deleted": true,
	})
}
