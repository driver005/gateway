package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

type FilterableInvite struct {
	core.FilterModel
}

type ListInvite struct {
	*models.Invite
	Token string `json:"token"`
}

// @oas:schema:AdminPostInvitesReq
// type: object
// required:
//   - user
//   - role
//
// properties:
//
//	user:
//	  description: "The email associated with the invite. Once the invite is accepted, the email will be associated with the created user."
//	  type: string
//	  format: email
//	role:
//	  description: "The role of the user to be created. This does not actually change the privileges of the user that is eventually created."
//	  type: string
//	  enum: [admin, member, developer]
type CreateInviteInput struct {
	Email string          `json:"email"`
	Role  models.UserRole `json:"role"`
}

// @oas:schema:AdminPostInvitesInviteAcceptReq
// type: object
// description: "The details of the invite to be accepted."
// required:
//   - token
//   - user
//
// properties:
//
//	token:
//	  description: "The token of the invite to accept. This is a unique token generated when the invite was created or resent."
//	  type: string
//	user:
//	  description: "The details of the user to create."
//	  type: object
//	  required:
//	    - first_name
//	    - last_name
//	    - password
//	  properties:
//	    first_name:
//	      type: string
//	      description: the first name of the User
//	    last_name:
//	      type: string
//	      description: the last name of the User
//	    password:
//	      description: The password for the User
//	      type: string
//	      format: password
type AcceptInvite struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}
