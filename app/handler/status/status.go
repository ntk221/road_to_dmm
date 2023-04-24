package status

import (
	"encoding/json"
	"log"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

type Attachment struct {
	ID          int64  `json:"id"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type StatusWithAccount struct {
	ID               int64           `json:"id"`
	Account          *object.Account `json:"account"`
	Content          string          `json:"content"`
	CreatedAt        object.DateTime `json:"created_at"`
	MediaAttachments []Attachment    `json:"media_attachments"`
}

// handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザー認証が成功した時の処理
	// Authentication Header から username を取得する
	/*a := r.Header.Get("Authorization")
	log.Printf("Authorization: %s", a)
	pair := strings.SplitN(a, " ", 2)
	if len(pair) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authType := pair[0]
	if !strings.EqualFold(authType, "username") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	username := pair[1]

	// username から account を取得する
	if account_rep := h.app.Dao.Account(); account_rep != nil {
		if account, err := account_rep.FindByUsername(ctx, username); err != nil {
			log.Printf("failed to find account: %v", err)
			httperror.InternalServerError(w, err)
			return
		} else if account == nil {
			log.Printf("account not found: %v", err)
			httperror.InternalServerError(w, err)
			return
		} else {
			log.Printf("account found: %v", account)
		}
	}*/

	status := new(object.Status)
	status.AccountID = 1 // TODO: get account id from context

	var statusMap map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&statusMap)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	if statusContent, ok := statusMap["status"].(string); ok {
		status.Content = statusContent
	}

	if status_rep := h.app.Dao.Status(); status_rep != nil {
		if err := status_rep.Create(ctx, status); err != nil {
			log.Printf("failed to create status: %v", err)
			httperror.InternalServerError(w, err)
			return
		}
	}

	if account_rep := h.app.Dao.Account(); account_rep != nil {
		if account, err := account_rep.FindByAccountID(ctx, status.AccountID); err != nil {
			log.Printf("failed to find account: %v", err)
			httperror.InternalServerError(w, err)
			return
		} else if account == nil {
			log.Printf("account not found: %v", err)
			httperror.InternalServerError(w, err)
			return
		} else {
			StatusWithAccount := &StatusWithAccount{
				ID:               status.ID,
				Account:          account,
				Content:          status.Content,
				MediaAttachments: []Attachment{},
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(StatusWithAccount); err != nil {
				log.Printf("failed to encode status: %v", err)
				httperror.InternalServerError(w, err)
				return
			}
		}
	}
}
