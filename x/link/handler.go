package link

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	cbd "github.com/cybercongress/cyberd/types"
	acc "github.com/cybercongress/cyberd/x/acc/types"
	"github.com/cybercongress/cyberd/x/link/keeper"
	cbdlink "github.com/cybercongress/cyberd/x/link/types"
)

// NewHandler returns a handler for "link" type messages.
// cis  - cids index storage
// ils  - links storage
// as   - account storage
// imms - in-memory storage
func NewLinksHandler(cis keeper.CidNumberKeeper, ls *keeper.LinkIndexedKeeper, as acc.AccountIndexKeeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {

		linkMsg := msg.(Msg)

		//validations
		//todo: optimize
		for _, link := range linkMsg.Links {
			// if cid not exists it automatically means that this is new link
			fromCidNumber, exists := cis.GetCidNumber(ctx, link.From)
			if !exists {
				continue
			}
			toCidNumber, exists := cis.GetCidNumber(ctx, link.To)
			if !exists {
				continue
			}

			accNumber := acc.AccNumber(as.GetAccount(ctx, linkMsg.Address).GetAccountNumber())
			compactLink := cbdlink.NewLink(fromCidNumber, toCidNumber, accNumber)

			if ls.IsLinkExist(compactLink) {
				return sdk.Result{Code: cbd.CodeLinkAlreadyExist, Codespace: cbd.CodespaceCbd}
			}
		}

		for _, link := range linkMsg.Links {
			fromCidNumber := cis.GetOrPutCidNumber(ctx, link.From)
			toCidNumber := cis.GetOrPutCidNumber(ctx, link.To)
			a := as.GetAccount(ctx, linkMsg.Address)
			as.AddToIndex(a) // TODO: Update with new version SDK
			accNumber := acc.AccNumber(a.GetAccountNumber())

			ls.PutLink(ctx, cbdlink.NewLink(fromCidNumber, toCidNumber, accNumber))
		}

		return sdk.Result{Code: cbd.CodeOK, Codespace: cbd.CodespaceCbd}
	}
}
