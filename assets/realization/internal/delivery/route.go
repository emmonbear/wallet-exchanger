package delivery

import (
	"github.com/gofiber/fiber/v2"
	//"gitlab.axarea.ru/main/aifory/svc/gateway/internal/middleware"
)

func MapClientV1Routes(route fiber.Router, h *Handler) {
	route = route.Group("v1")

	//route.Get("/fetch_exchange_currencies", mw.ClientAuth, h.FetchExchangeCurrencies)
	//
	//route.Post("/solve_captcha", h.SolveCaptcha)
	//
	//route.Post("/locations", mw.ClientAuth, h.FetchLocations)
	//
	//route.Get("/invited", mw.ClientAuth, h.GetNumOfInvitedUsers)

	client := route.Group("client")
	{
		client.Get("/balance", h.GetForeignCurrencyCardpayment)
		client.Get("/balance_localization", h.GetForeignCurrencyCardpayment)
		//userSettingsGroup := client.Group("user_settings")
		//userSettingsGroup.Post("/update", mw.ClientAuth, h.UpdateUserSettings)
		//userSettingsGroup.Post("/get_invite_link", mw.ClientAuth, h.GetInviteLink)
		//userSettingsGroup.Post("/set_user_initials", mw.ClientAuth, h.SetUserInitials)
		//client.Post("/send_tg_code", mw.ClientAuth, h.SendConfirmationCode)
	}

}

//func MapClientV1Routes(route fiber.Router, mw middleware.Middleware, h *Handler) {
//	route = route.Group("v1")
//
//	route.Get("/fetch_exchange_currencies", mw.ClientAuth, h.FetchExchangeCurrencies)
//
//	route.Post("/solve_captcha", h.SolveCaptcha)
//
//	route.Post("/locations", mw.ClientAuth, h.FetchLocations)
//
//	route.Get("/invited", mw.ClientAuth, h.GetNumOfInvitedUsers)
//
//	files := route.Group("files")
//	{
//		// list of downloaded files
//		files.Post("/list", h.GetUploadedFilesList)
//
//		files.Post("/download/:filename", h.DownloadFile)
//
//		// upload.
//		files.Post("/excel", h.ParseExcelFile)
//		files.Post("/save/db", h.SaveFileInDatabase)
//		files.Post("/save/s3", h.SaveFileInS3)
//
//		files.Get("/pentest/main", h.GetFilesHTML)
//	}
//
//	emailGroup := route.Group("email")
//	{
//		emailGroup.Post("/create", mw.ClientAuth, h.SendEmailConfirmationCodeRequest)
//		emailGroup.Post("/delete", mw.ClientAuth, h.DeleteEmail)
//		emailGroup.Post("/confirm", h.SetUserEmail)
//		emailGroup.Post("/reset_password", h.SendPasswordResetCode)
//		emailGroup.Post("/get_reset_password_token", h.GetEmailPasswordToken)
//	}
//	currencyRoute := route.Group("currency")
//	{
//		currencyRoute.Get("/", mw.ClientAuth, h.FetchCurrencyList)
//		currencyRoute.Post("/locale", mw.ClientAuth, h.FetchCurrencyListByLocale)
//	}
//	swapRoute := route.Group("swap")
//	{
//		swapRoute.Get("/limits", mw.ClientAuth, h.FetchSwapLimits)
//	}
//
//	accountsGroup := route.Group("accounts")
//	{
//		accountsGroup.Post("/", mw.ClientAuth, h.FetchAccountList)
//		accountsGroup.Group("create").Post("/crypto", mw.ClientAuth, h.CreateCryptoAccount)
//		accountsGroup.Group("create").Post("/fiat", mw.ClientAuth, h.CreateFiatAccount)
//		accountsGroup.Post("/update_wallet", mw.ClientAuth, h.UpdateWallet)
//		accountsGroup.Post("/delete_wallet", mw.ClientAuth, h.DeleteWallet)
//		accountsGroup.Post("/check-balance", mw.ClientAuth, h.CheckBalance)
//	}
//
//	withdrawGroup := route.Group("withdraw")
//	{
//		crypto := withdrawGroup.Group("crypto")
//		crypto.Post("/", mw.ClientAuth, mw.VerifyConfirmationCode, h.WithdrawCrypto)
//
//		fiat := withdrawGroup.Group("fiat")
//		fiat.Post("/", mw.ClientAuth, mw.VerifyConfirmationCode, h.WithdrawFiat)
//		fiat.Post("/sbp", mw.ClientAuth, mw.VerifyConfirmationCode, h.WithdrawFiatSBP)
//		fiat.Post("/sbp-methods", mw.ClientAuth, h.FetchWithdrawSBPMethods)
//		//fiat.Post("/sbp", mw.ClientAuth, mw.VerifyConfirmationCode, h.WithdrawFiatSBP)
//		//fiat.Post("/sbp-methods", mw.ClientAuth, h.FetchWithdrawSBPMethods)
//		cardpayment := withdrawGroup.Group("cardpayment")
//		cardpayment.Post("/", mw.ClientAuth, h.WithdrawCardpayment)
//	}
//
//	tx := route.Group("tx")
//	{
//		tx.Post("/", mw.ClientAuth, h.FetchTxByByAccount)
//	}
//
//	details := route.Group("details")
//	{
//		details.Post("/user", mw.ClientAuth, h.GetUserDetails)
//		details.Post("/order/info", mw.ClientAuth, h.FetchOrderByOrderID)
//		details.Post("/numberOrders", mw.ClientAuth, h.FetchNumbersOfOrders)
//		details.Post("/allorders", mw.ClientAuth, h.FetchAllOrders)
//	}
//
//	invoiceFiat := route.Group("invoice")
//	{
//		fiat := invoiceFiat.Group("fiat")
//		{
//			fiat.Post("/p2p", mw.ClientAuth, h.InvoiceFiatP2P)
//			fiat.Post("/h2h", mw.ClientAuth, h.InvoiceFiatH2H)
//			fiat.Post("/p2p-methods", mw.ClientAuth, h.FetchP2PMethodsList)
//			fiat.Post("/p2p/cancel", mw.ClientAuth, h.CancelP2POrder)
//			fiat.Post("/p2p/confirm", mw.ClientAuth, h.ConfirmP2POrder)
//			fiat.Post("/sbp", mw.ClientAuth, h.InvoiceFiatSBP)
//			fiat.Post("/sbp-methods", mw.ClientAuth, h.FetchSBPMethodsList)
//			fiat.Post("/sbp/confirm", mw.ClientAuth, h.ConfirmSBPOrder)
//		}
//	}
//
//	fiatGroup := route.Group("fiat")
//	{
//		certGroup := fiatGroup.Group("cert")
//		{
//			certGroup.Post("/create", mw.ClientAuth, mw.VerifyConfirmationCode, h.FiatTransferCertCreate)
//			certGroup.Post("/use", mw.ClientAuth, mw.CertAntiFraud, h.FiatTransferCertUse)
//			certGroup.Post("/cancel", mw.ClientAuth, mw.CertAntiFraud, h.FiatTransferCertCancel)
//			certGroup.Post("/order/cancel", mw.ClientAuth, h.CancelOrder)
//			certGroup.Post("/preview", mw.ClientAuth, mw.CertAntiFraud, h.FiatTransferCertPreview)
//		}
//	}
//
//	cardpaymentGroup := route.Group("cardpayment")
//	{
//		certGroup := cardpaymentGroup.Group("currency")
//		certGroup.Get("/get_foreign_currency", mw.ClientAuth, h.GetForeignCurrencyCardpayment)
//	}
//
//	route.Post("/order/cancel", mw.ClientAuth, h.CancelOrder)
//
//	// WARNING WEBHOOK - DANGEROUS ACTIONS
//	webhook := route.Group("webhook")
//	{
//		webhook.Post("/invoice", mw.Signature, h.WebhookInvoice)
//		webhook.Post("/withdraw", mw.Signature, h.WebhookWithdraw)
//		webhook.Post("/withdrawcardpayment", mw.Signature, h.WebhookWithdrawCardpayment)
//	}
//
//	exchange := route.Group("exchange")
//	{
//		// buy/sell crypto are deprecated
//		//exchange.Post("/sell_crypto", mw.ClientAuth, h.SellCrypto)
//		//exchange.Post("/buy_crypto", mw.ClientAuth, h.BuyCrypto)
//		exchange.Post("/", mw.ClientAuth, h.ClientExchange)
//
//		exchange.Post("list-transactions", mw.ClientAuth, h.ClientListTransactions)
//	}
//
//	auth := route.Group("auth")
//	{
//		clientGroup := auth.Group("client")
//		{
//			mobileAuth := clientGroup.Group("/mobile")
//			{
//				mobileAuth.Post("/sign-in/request", mw.BruteForceMW, h.MobileClientSignIn)
//				mobileAuth.Post("/sign-in/confirm", mw.BruteForceMW, h.MobileClientSignInConfirm)
//				mobileAuth.Post("/websocket-token", mw.ClientAuth, h.GetWebSocketAuthMobileToken)
//
//				signUpMobileGroup := mobileAuth.Group("sign-up")
//				{
//					signUpMobileGroup.Post("/request", h.MobileClientSignUpRequest)
//					signUpMobileGroup.Post("/email/confirm", h.MobileClientConfirmLinkEmailOnSignUp)
//					signUpMobileGroup.Post("/email/create", h.MobileClientLinkEmailOnSignUp)
//					signUpMobileGroup.Post("/tg/create", h.MobileClientLinkTgOnSignUp)
//					signUpMobileGroup.Post("/tg/check_link", h.MobileClientCheckLinkTgOnSignUp)
//					signUpMobileGroup.Post("/finalize", h.MobileClientFinalizeSignUp)
//				}
//			}
//
//			signUpGroup := clientGroup.Group("sign-up")
//			{
//				signUpGroup.Post("/request", h.ClientSignUpRequest)
//				signUpGroup.Post("/email/confirm", h.ClientConfirmLinkEmailOnSignUp)
//				signUpGroup.Post("/email/create", h.ClientLinkEmailOnSignUp)
//				signUpGroup.Post("/tg/create", h.ClientLinkTgOnSignUp)
//				signUpGroup.Post("/tg/check_link", h.ClientCheckLinkTgOnSignUp)
//				signUpGroup.Post("/finalize", h.ClientFinalizeSignUp)
//			}
//
//			clientGroup.Post("/sign-in/request", mw.BruteForceMW, h.AuthClientSignInRequest)
//			clientGroup.Post("/sign-in/confirm", mw.BruteForceMW, h.AuthClientSignInConfirm)
//			clientGroup.Post("/sign-in/finalize", mw.BruteForceMW, h.AuthClientSignInFinalize)
//			clientGroup.Post("/sign-out", h.AuthClientSignOut)
//			clientGroup.Post("/renew", mw.BruteForceMW, h.AuthClientRenew)
//			clientGroup.Post("/reset_pin", h.ResetPin)
//
//			clientGroup.Post("/get_limit", mw.ClientAuth, h.GetLimit)
//			clientGroup.Post("/get_limits", h.GetLimits)
//			clientGroup.Post("/set_limits", mw.ClientAuth, h.SetLimits)
//			clientGroup.Post("/verify_kyc", mw.SignatureKYC, h.AuthVerifyKYC)
//			clientGroup.Post("/get_kyc_verify_link", mw.ClientAuth, h.GetKUCVerificationLink)
//
//			clientGroup.Post("/change_telegram", mw.BruteForceMW, mw.ClientAuth, mw.VerifyConfirmationCode, h.UpdateTelegram)
//			clientGroup.Post("/new_telegram", mw.BruteForceMW, mw.ClientAuth, h.SetNewTelegram)
//
//			clientGroup.Post("/change_password", h.ChangePassword)
//
//			clientGroup.Get("/otp", mw.BruteForceMW, mw.ClientAuth, h.ChangeOTP)
//			clientGroup.Get("/otp_mobile", mw.BruteForceMW, mw.ClientAuth, h.ChangeOTPMobile)
//			clientGroup.Post("/otp/confirm", mw.BruteForceMW, mw.ClientAuth, h.ChangeOTPConfirm)
//			clientGroup.Post("/otp/drop", mw.ClientAuth, mw.VerifyConfirmationCode, h.ClientDropOTP)
//
//			clientGroup.Get("/delete", mw.ClientAuth, h.SoftDeleteUser)
//			clientGroup.Get("/undo_delete", mw.ClientAuth, h.UndoDeleteUser)
//
//			clientGroup.Post("recovery_phrase", h.VerifyByRecoveryPhrase)
//
//			clientGroup.Post("websocket-token", mw.ClientAuth, h.GetWebSocketAuthClientToken)
//
//		}
//	}
//
//	cash := route.Group("cash")
//	{
//		cash.Post("/invoice", mw.ClientAuth, h.MakeCashInvoiceOrder)
//		cash.Post("/withdraw", mw.ClientAuth, h.MakeCashWithdrawOrder)
//		cash.Post("/withdraw/pro", mw.ClientAuth, h.MakeCashWithdrawOrderPro)
//		cash.Post("/withdraw/entries", mw.ClientAuth, h.PrepareWithdrawEntries)
//		cash.Post("/office/getall", mw.ClientAuth, h.GetAllOffices)
//		cash.Post("/office/fetch-slots", mw.ClientAuth, h.FetchTimeSlots)
//		cash.Post("/order/cancel", mw.ClientAuth, h.CancelOrder)
//	}
//
//	tariff := route.Group("tariff")
//	{
//		calculate := tariff.Group("calculate")
//		calculate.Post("/sell", mw.ClientAuth, h.CalculateExchangeRatesSell)
//		calculate.Post("/buy", mw.ClientAuth, h.CalculateExchangeRatesBuy)
//
//		invoice := tariff.Group("invoice")
//		invoice.Post("/crypto", mw.ClientAuth, h.GetCryptoInvoiceTariff)
//		invoice.Post("/fiat", mw.ClientAuth, h.GetFiatInvoiceTariff)
//		invoice.Post("/cash", mw.ClientAuth, h.GetCashInvoiceTariff)
//		invoice.Post("/sbp", mw.ClientAuth, h.GetSBPInvoiceTariff)
//
//		withdraw := tariff.Group("withdraw")
//		withdraw.Post("/crypto", mw.ClientAuth, h.GetCryptoWithdrawTariff)
//		withdraw.Post("/fiat", mw.ClientAuth, h.GetFiatWithdrawTariff)
//		withdraw.Post("/cardpayment", mw.ClientAuth, h.GetCardpaymentWithdrawTariff)
//		withdraw.Post("/cash", mw.ClientAuth, h.GetCashWithdrawTariff)
//		withdraw.Post("/sbp", mw.ClientAuth, h.GetSBPWithdrawTariff)
//	}
//
//	client := route.Group("client")
//	{
//		client.Get("/balance_cold_config", h.GetBalanceColdConfig)
//		client.Get("/balance_localization", h.GetBalanceLocalization)
//		userSettingsGroup := client.Group("user_settings")
//		userSettingsGroup.Post("/update", mw.ClientAuth, h.UpdateUserSettings)
//		userSettingsGroup.Post("/get_invite_link", mw.ClientAuth, h.GetInviteLink)
//		userSettingsGroup.Post("/set_user_initials", mw.ClientAuth, h.SetUserInitials)
//		client.Post("/send_tg_code", mw.ClientAuth, h.SendConfirmationCode)
//	}
//
//	banner := route.Group("banner")
//	{
//		banner.Post("/get", mw.ClientAuth, h.GetBanner)
//	}
//
//	cbTransfer := route.Group("cb-transfer")
//	{
//		order := cbTransfer.Group("order")
//		{
//			order.Post("create", mw.ClientAuth, h.ClientCreateOrder)
//			order.Post("fetch", mw.ClientAuth, h.ClientFetchOrders)
//			order.Post("confirm", mw.ClientAuth, h.ClientConfirmOrder)
//			order.Post("cancel", mw.ClientAuth, h.ClientCancelOrder)
//			order.Post("change-account", mw.ClientAuth, h.ClientChangeAccount)
//		}
//
//		cbTransfer.Post("get-file", mw.ClientAuth, h.ClientDownloadFile)
//	}
//
//}

//func MapAdminV1Routes(route fiber.Router, mw middleware.Middleware, h *Handler) {
//	route = route.Group("v1")
//
//	route.Post("/locations", mw.SetTreasurerAccess, mw.Auth, h.FetchLocations)
//
//	exchange := route.Group("exchange")
//	{
//		exchangeGroup := exchange.Group("/admin")
//		{
//			exchangeGroup.Post("/list_transactions",
//				mw.SetTreasurerAccess,
//				mw.SetManagerAccess,
//				mw.SetElderManagerAccess,
//				mw.SetRaterAccess,
//				mw.SetDeveloperAccess,
//				mw.SetTreasurerAccess,
//				mw.Auth, h.AdminListTransactions)
//			exchangeGroup.Post("/fetch_orders",
//				mw.SetTreasurerAccess,
//				mw.SetManagerAccess,
//				mw.SetElderManagerAccess,
//				mw.SetRaterAccess,
//				mw.SetDeveloperAccess,
//				mw.SetTreasurerAccess,
//				mw.Auth, h.AdminFetchExchangeOrders)
//			exchangeGroup.Post("/list_operations",
//				mw.SetRaterAccess,
//				mw.SetElderManagerAccess,
//				mw.SetRaterAccess,
//				mw.SetDeveloperAccess,
//				mw.SetTreasurerAccess,
//				mw.Auth, h.AdminListOperations)
//			exchangeGroup.Post("/get_detailed_order_info",
//				mw.SetTreasurerAccess,
//				mw.SetManagerAccess,
//				mw.SetElderManagerAccess,
//				mw.SetRaterAccess,
//				mw.SetDeveloperAccess,
//				mw.SetTreasurerAccess,
//				mw.Auth, h.AdminGetDetailedOrderInfo)
//			exchangeGroup.Post("/create_operation", mw.SetTreasurerAccess, mw.Auth, h.AdminCreateOperation)
//			exchangeGroup.Post("/cancel_operation", mw.SetTreasurerAccess, mw.Auth, h.AdminCancelOperation)
//			exchangeGroup.Post("/calculate_margin", mw.SetTreasurerAccess, mw.Auth, h.AdminCalculatePreliminaryMargin)
//			exchangeGroup.Post("/lock_transactions", mw.SetTreasurerAccess, mw.Auth, h.AdminLockTransactions)
//			exchangeGroup.Post("/unlock_transactions", mw.SetTreasurerAccess, mw.Auth, h.AdminUnlockTransactions)
//			exchangeGroup.Post("/sell_crypto", mw.SetTreasurerAccess, mw.Auth, h.AdminSellCrypto)
//			exchangeGroup.Post("/buy_crypto", mw.SetTreasurerAccess, mw.Auth, h.AdminBuyCrypto)
//			exchangeGroup.Post("/statistic", mw.SetTreasurerAccess, mw.Auth, h.AdminStatistics)
//			exchangeGroup.Post("/plan", mw.SetTreasurerAccess, mw.Auth, h.AdminMonthlyPlan)
//			exchangeGroup.Post("/margin", mw.SetTreasurerAccess, mw.Auth, h.AdminForceMargin)
//			exchangeGroup.Post("/set_accepted_transaction", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.AdminSetAcceptedTransaction)
//		}
//	}
//
//	swap := route.Group("/swap")
//	{
//		swap.Post("/fetch",
//			mw.SetRaterAccess,
//			mw.SetElderManagerAccess,
//			mw.SetRaterAccess,
//			mw.SetDeveloperAccess,
//			mw.SetTreasurerAccess,
//			mw.Auth,
//			h.FetchSwapOrders,
//		)
//		swap.Post("/details",
//			mw.SetRaterAccess,
//			mw.SetElderManagerAccess,
//			mw.SetRaterAccess,
//			mw.SetDeveloperAccess,
//			mw.SetTreasurerAccess,
//			mw.Auth,
//			h.AdminSwapDetails,
//		)
//		swap.Post("/confirm",
//			mw.SetRaterAccess,
//			mw.SetElderManagerAccess,
//			mw.SetRaterAccess,
//			mw.SetDeveloperAccess,
//			mw.SetTreasurerAccess,
//			mw.Auth,
//			h.AdminSwapConfirm,
//		)
//		swap.Post("/cancel",
//			mw.SetRaterAccess,
//			mw.SetElderManagerAccess,
//			mw.SetRaterAccess,
//			mw.SetDeveloperAccess,
//			mw.SetTreasurerAccess,
//			mw.Auth,
//			h.AdminSwapCancel,
//		)
//	}
//
//	auth := route.Group("auth")
//	{
//		adminGroup := auth.Group("/admin")
//		{
//			adminGroup.Post("/sign-up/request", h.AuthAdminSignUpRequest)
//			adminGroup.Post("/sign-up/verify", h.AuthAdminSignUpVerify)
//			adminGroup.Post("/sign-in", h.AuthAdminSignIn)
//			adminGroup.Post("/sign-in/nal", h.AdminNalSignIn)
//			adminGroup.Post("/invite/admin", mw.SetTreasurerAccess, mw.Auth, h.CreateAdminInviteToken)
//			adminGroup.Post("/invite/client",
//				mw.SetElderManagerAccess, mw.SetTreasurerAccess, mw.Auth,
//				h.CreateClientInviteToken)
//			adminGroup.Get("/invite/origins",
//				mw.SetElderManagerAccess, mw.SetTreasurerAccess, mw.Auth,
//				h.GetClientInviteTokenOrigins)
//			adminGroup.Post("/sign-out", h.AuthAdminSignOut)
//			adminGroup.Post("/update-user",
//				mw.SetElderManagerAccess,
//				mw.SetTreasurerAccess, mw.Auth, h.UpdateUserInfo)
//
//			adminGroup.Post("/fetch-users",
//				mw.SetDeveloperAccess,
//				mw.SetRaterAccess,
//				mw.SetManagerAccess,
//				mw.SetElderManagerAccess,
//				mw.SetTreasurerAccess, mw.Auth, h.FetchUsersWithTariff)
//
//			adminGroup.Post("websocket-token", mw.Auth, h.GetWebSocketAuthAdminToken)
//		}
//	}
//
//	tariff := route.Group("tariff")
//	{
//		adminGroup := tariff.Group("admin")
//		{
//			adminGroup.Post("/add-client",
//				mw.SetRaterAccess,
//				mw.SetTreasurerAccess, mw.Auth, h.AddClient)
//
//			adminGroup.Post("/get-all",
//				mw.SetDeveloperAccess,
//				mw.SetElderManagerAccess,
//				mw.SetRaterAccess,
//				mw.SetTreasurerAccess, mw.Auth, h.GetAllClientsWithTariffs)
//
//			adminGroup.Post("/change-tariff",
//				mw.SetRaterAccess,
//				mw.SetTreasurerAccess, mw.Auth, h.ChangeClientTariffType)
//
//			adminGroup.Post("/calculate-spread",
//				mw.SetRaterAccess,
//				mw.SetTreasurerAccess, mw.Auth, h.CalculateSpread)
//
//			adminGroup.Post("/freeze-rate",
//				mw.SetRaterAccess,
//				mw.SetTreasurerAccess, mw.Auth, h.FreezeRate)
//
//			adminGroup.Post("/unfreeze-rate",
//				mw.SetRaterAccess,
//				mw.SetTreasurerAccess, mw.Auth, h.UnfreezeRate)
//
//			baseTariffGroup := adminGroup.Group("base-tariff")
//			{
//				baseTariffGroup.Post("/get-list",
//					mw.SetDeveloperAccess,
//					mw.SetManagerAccess,
//					mw.SetElderManagerAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess, mw.Auth, h.GetBaseTariffList)
//
//				baseTariffGroup.Post("/add",
//					mw.SetElderManagerAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess, mw.Auth, h.AddBaseTariff)
//
//				baseTariffGroup.Post("/duplicate",
//					mw.SetElderManagerAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess, mw.Auth,
//					h.DublicateBaseTariff)
//
//				baseTariffGroup.Post("/get/:id",
//					mw.SetDeveloperAccess,
//					mw.SetManagerAccess,
//					mw.SetElderManagerAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess, mw.Auth, h.GetBaseTariff)
//
//				baseTariffGroup.Post("/update",
//					mw.SetElderManagerAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess, mw.Auth, h.UpdateBaseTariff)
//
//				baseTariffGroup.Post("/set",
//					mw.SetElderManagerAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess, mw.Auth, h.SetClientBaseTariff)
//			}
//			flexTariffGroup := adminGroup.Group("flex-tariff")
//			{
//				flexTariffGroup.Post("/get",
//					mw.SetDeveloperAccess,
//					mw.SetRaterAccess,
//					mw.SetManagerAccess,
//					mw.SetElderManagerAccess,
//					mw.SetTreasurerAccess, mw.Auth, h.GetFlexTariff)
//
//				flexTariffGroup.Post("/set",
//					mw.SetElderManagerAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess, mw.Auth, h.SetClientFlexTariff)
//			}
//		}
//	}
//
//	admin := route.Group("admin")
//	{
//		admin.Post("/ban", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.BanUsers)
//		admin.Post("/unban", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.UnbanUsers)
//		admin.Post("/ban_employee", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.BanEmployee)
//		admin.Post("/unban_employee", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.UnbanEmployee)
//		admin.Post("/drop_telegram", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.DropTelegram)
//		admin.Post("/drop_otp", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.AdminDropOTPForClient)
//		admin.Post("/fetch_employees",
//			mw.SetManagerAccess,
//			mw.SetElderManagerAccess,
//			mw.SetDeveloperAccess,
//			mw.SetTreasurerAccess,
//			mw.Auth, h.FetchEmployees)
//		admin.Post("/fetch_roles", mw.SetAllRolesAccess, mw.Auth, h.FetchRoles)
//		admin.Post("/fetch_admin_orders", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.FetchAdminOrders)
//		admin.Post("/generate-address", mw.SetTreasurerAccess, mw.Auth, h.GenerateClientAddress)
//
//		admin.Post("/fetch_roles_test", h.FetchRoles)
//
//		admin.Post("/get_invited_users_list", mw.SetElderManagerAccess, mw.Auth, h.GetInvitedUsersList)
//
//		admin.Get("/fetch_clients_balances", mw.Auth, h.FetchClientsBalances)
//
//		admin.Post("/fetch_client_orders", mw.SetElderManagerAccess, mw.Auth, h.FetchClientOrders)
//
//		bannerGroup := admin.Group("banner")
//		{
//			bannerGroup.Post("/set_status", mw.Auth, h.SetBannerStatus)
//		}
//
//		changeGroup := admin.Group("change")
//		{
//			changeGroup.Get("/otp", mw.SetTreasurerAccess, mw.Auth, h.ChangeOTP)
//			changeGroup.Post("/otp/confirm", mw.SetTreasurerAccess, mw.Auth, h.ChangeOTPConfirm)
//
//			changeGroup.Post("/password", mw.SetTreasurerAccess, mw.Auth, h.ChangePassword)
//		}
//		admin.Get("/details/user", mw.SetAllRolesAccess, mw.Auth, h.GetEmployeeDetails)
//		admin.Get("/client-accounts",
//			mw.SetElderManagerAccess,
//			mw.SetDeveloperAccess,
//			mw.SetTreasurerAccess,
//			mw.Auth, h.AdminGetClientAccounts)
//		admin.Get("/accounts-types", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminGetAccountsTypes)
//		admin.Get("/accounts-types/:transactionType",
//			mw.SetDeveloperAccess,
//			mw.SetTreasurerAccess, mw.Auth, h.AdminGetAccountsTypes)
//
//		admin.Post("/fiat/order",
//			mw.SetRaterAccess,
//			mw.SetDeveloperAccess,
//			mw.SetManagerAccess,
//			mw.SetElderManagerAccess,
//			mw.SetTreasurerAccess, mw.Auth, h.FetchFiatOrders)
//
//		cryptoGroup := admin.Group("/crypto")
//		{
//			orderGroup := cryptoGroup.Group("/order")
//			{
//				orderGroup.Post("/",
//					mw.SetRaterAccess,
//					mw.SetDeveloperAccess,
//					mw.SetManagerAccess,
//					mw.SetElderManagerAccess,
//					mw.SetTreasurerAccess, mw.Auth, h.FetchCryptoOrders)
//				orderGroup.Post("/invoice/confirm", mw.SetTreasurerAccess, mw.Auth, h.ConfirmCryptoInvoice)
//				orderGroup.Post("/withdraw/confirm",
//					mw.SetTreasurerAccess,
//					mw.Auth,
//					mw.VerifyConfirmationCode,
//					h.ConfirmCryptoWithdraw)
//			}
//
//			modeGroup := cryptoGroup.Group("/mode")
//			{
//				modeGroup.Post("/set", mw.SetTreasurerAccess, mw.Auth, h.SetConfirmationMode)
//				modeGroup.Post("/get", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.GetConfirmationMode)
//			}
//
//			currencyGroup := cryptoGroup.Group("/currency")
//			{
//				currencyGroup.Post("/get", mw.SetDeveloperAccess, mw.SetRaterAccess, mw.SetTreasurerAccess, mw.Auth, h.FetchCurrencies)
//				currencyGroup.Post("/update", mw.SetTreasurerAccess, mw.Auth, h.UpdateCurrency)
//			}
//
//			walletGroup := cryptoGroup.Group("/wallet")
//			{
//				clientGroup := walletGroup.Group("/client")
//				{
//					clientGroup.Post("/get", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.GetClientWallets)
//					clientGroup.Post("/withdraw", mw.SetTreasurerAccess, mw.Auth, h.MakeClientWithdrawal)
//
//				}
//				hotGroup := walletGroup.Group("/hot")
//				{
//					hotGroup.Post("/get", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.GetHotWallets)
//					hotGroup.Post("/set", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.SetHotWalletStatus)
//					hotGroup.Post("/withdraw", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.MakeHotWithdrawal)
//					hotGroup.Post("/add", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.AddHotWallet)
//
//				}
//
//				cashGroup := walletGroup.Group("/cash")
//				{
//					cashGroup.Post("/get",
//						mw.SetDeveloperAccess,
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						h.GetCashWallets)
//					cashGroup.Post("/withdraw",
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						mw.VerifyConfirmationCode,
//						h.MakeCashWithdrawal)
//					cashGroup.Post("/add",
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						mw.VerifyConfirmationCode,
//						h.AddCashWallet)
//					cashGroup.Post("/delete",
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						mw.VerifyConfirmationCode,
//						h.DeleteCashWallet)
//					cashGroup.Post("/change",
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						mw.VerifyConfirmationCode,
//						h.ChangeCashWallet)
//				}
//
//				bourseGroup := walletGroup.Group("/bourse")
//				{
//					bourseGroup.Post("/get",
//						mw.SetDeveloperAccess,
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						h.GetBourseWallets)
//					bourseGroup.Post("/withdraw",
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						mw.VerifyConfirmationCode,
//						h.MakeBourseWithdrawal)
//					bourseGroup.Post("/add",
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						mw.VerifyConfirmationCode,
//						h.AddBourseWallet)
//					bourseGroup.Post("/delete",
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						mw.VerifyConfirmationCode,
//						h.DeleteBourseWallet)
//					bourseGroup.Post("/change",
//						mw.SetTreasurerAccess,
//						mw.SetNalAccess,
//						mw.AuthNal,
//						mw.VerifyConfirmationCode,
//						h.ChangeBourseWallet)
//				}
//
//				coldGroup := walletGroup.Group("/cold")
//				{
//					coldGroup.Post("/get", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.GetColdWallets)
//					coldGroup.Post("/add", mw.SetTreasurerAccess, mw.Auth, h.AddColdWallet)
//					coldGroup.Post("/set-status", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.SetColdWalletStatus)
//
//				}
//
//				adminGroup := walletGroup.Group("/admin")
//				{
//					adminGroup.Post("/get", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.GetAdminWallets)
//					adminGroup.Post("/add", mw.SetTreasurerAccess, mw.Auth, h.AddAdminWallet)
//
//				}
//			}
//		}
//
//		admin.Post("/fiat/order/fetch",
//			mw.SetCashierAccess,
//			mw.SetElderCashierAccess,
//			mw.SetDeveloperAccess,
//			mw.SetRaterAccess,
//			mw.SetManagerAccess,
//			mw.SetElderManagerAccess,
//			mw.SetTreasurerAccess,
//			mw.Auth, h.FetchFiatOrder)
//		admin.Post("/crypto/order/fetch",
//			mw.SetDeveloperAccess,
//			mw.SetRaterAccess,
//			mw.SetManagerAccess,
//			mw.SetElderManagerAccess,
//			mw.SetTreasurerAccess,
//			mw.Auth, h.FetchCryptoOrder)
//
//		admin.Get("/currencies",
//			mw.SetManagerAccess,
//			mw.SetElderManagerAccess,
//			mw.SetDeveloperAccess,
//			mw.SetRaterAccess,
//			mw.SetTreasurerAccess,
//			mw.Auth, h.AdminGetCurrencies)
//		admin.Post("/similar_currencies", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.SetNalAccess, mw.Auth, h.GetSimilarBipTypeCurrencies)
//		admin.Post("/get_address_balance", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.SetNalAccess, mw.Auth, h.GetAddressBalance)
//
//		treasuryGroup := admin.Group("/treasury")
//		{
//			treasuryGroup.Post("/", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryFetchSummary)
//			treasuryGroup.Post("/delta", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryFetchSummaryDelta)
//
//			accountsGroup := treasuryGroup.Group("/accounts")
//			{
//				accountsGroup.Post("/passive-grouped", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryFetchPassiveAccountsGrouped)
//				accountsGroup.Post("/active-grouped", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryFetchActiveAccountsGrouped)
//				accountsGroup.Post("/passive", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryFetchPassiveAccounts)
//				accountsGroup.Post("/active", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryFetchActiveAccounts)
//			}
//
//			transactionsGroup := treasuryGroup.Group("/transactions")
//			{
//				transactionsGroup.Post("/details", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.FetchInternalTransactionDetails)
//				transactionsGroup.Post("/", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryGetInternalTransactions)
//				transactionsGroup.Post("/accounts/:transactionType", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryGetAccountsForInternalTransactions)
//				transactionsGroup.Post("/new", mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryCreateInternalTransaction)
//				transactionsGroup.Post("/close", mw.SetTreasurerAccess, mw.Auth, mw.VerifyConfirmationCode, h.AdminTreasuryCloseInternalTransaction)
//				transactionsGroup.Get("/types", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryGetInternalTransactionsTypes)
//				transactionsGroup.Get("/types/expenses", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.AdminTreasuryGetExpensesTypes)
//			}
//		}
//
//		settingsGroup := admin.Group("/settings")
//		{
//			settingsGroup.Get("/get_waf_rule_status", mw.SetTreasurerAccess, mw.Auth, h.GetWafRuleStatus)
//			settingsGroup.Post("/switch_waf_rule", mw.SetTreasurerAccess, mw.Auth, h.SwitchWafRule)
//		}
//
//		swapGroup := admin.Group("swap")
//		{
//			swapGroup.Post("/order",
//				mw.SetRaterAccess,
//				mw.SetDeveloperAccess,
//				mw.SetManagerAccess,
//				mw.SetElderManagerAccess,
//				mw.SetTreasurerAccess, mw.Auth, h.FetchSwapOrders)
//		}
//
//	}
//
//	cash := route.Group("cash")
//	{
//		adminRoutes := cash.Group("admin")
//		{
//			orderGroup := adminRoutes.Group("order")
//			{
//				orderGroup.Post("/update",
//					mw.SetCashierAccess,
//					mw.SetElderCashierAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess,
//					mw.Auth, h.UpdateCashOrder)
//				orderGroup.Post("/prepare-confirm",
//					mw.SetCashierAccess,
//					mw.SetElderCashierAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess,
//					mw.Auth, h.PrepareConfirmOrder)
//
//				orderGroup.Post("/fetch", mw.SetAllRolesAccess, mw.Auth, h.FetchCashOrders)
//
//				orderGroup.Post("/withdraw/confirm",
//					mw.SetCashierAccess,
//					mw.SetElderCashierAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess,
//					mw.Auth, h.FinalizeConfirmWithdrawOrder)
//
//				orderGroup.Put("/withdraw/update",
//					mw.SetCashierAccess,
//					mw.SetElderCashierAccess,
//					mw.SetRaterAccess,
//					mw.SetTreasurerAccess,
//					mw.Auth, h.UpdateWithdrawOrderStatus)
//
//				orderGroup.Post("/withdraw/fetch",
//					mw.SetManagerAccess,
//					mw.SetElderManagerAccess,
//					mw.SetRaterAccess,
//					mw.SetCashierAccess,
//					mw.SetRaterAccess,
//					mw.SetElderCashierAccess,
//					mw.SetTreasurerAccess,
//					mw.Auth, h.FetchWithdrawOrderByID)
//
//				orderGroup.Post("/invoice/fetch",
//					mw.SetManagerAccess,
//					mw.SetElderManagerAccess,
//					mw.SetCashierAccess,
//					mw.SetRaterAccess,
//					mw.SetElderCashierAccess,
//					mw.SetTreasurerAccess,
//					mw.Auth, h.FetchInvoiceOrderByID)
//				orderGroup.Put("/invoice/update",
//					mw.SetRaterAccess,
//					mw.SetCashierAccess,
//					mw.SetElderCashierAccess,
//					mw.SetTreasurerAccess,
//					mw.Auth, h.UpdateInvoiceOrderStatus)
//				orderGroup.Post("/invoice/confirm",
//					mw.SetRaterAccess,
//					mw.SetCashierAccess,
//					mw.SetElderCashierAccess,
//					mw.SetTreasurerAccess,
//					mw.Auth, h.FinalizeConfirmInvoiceOrder)
//
//			}
//
//			officeGroup := adminRoutes.Group("office")
//			{
//				officeGroup.Post("/getall", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.GetAllOffices)
//				officeGroup.Put("/update", mw.SetTreasurerAccess, mw.Auth, h.UpdateOffice)
//				officeGroup.Post("/create", mw.SetTreasurerAccess, mw.Auth, h.CreateOffice)
//				officeGroup.Post("/get", mw.SetDeveloperAccess, mw.SetTreasurerAccess, mw.Auth, h.GetOfficeByID)
//				officeGroup.Delete("/delete", mw.SetTreasurerAccess, mw.Auth, h.DeleteOffice)
//			}
//		}
//	}
//
//	cbTransfer := route.Group("cb-transfer")
//	{
//		order := cbTransfer.Group("order")
//		{
//			order.Post("fetch", mw.Auth, h.AdminFetchOrders)
//			order.Post("add-sum", mw.Auth, h.AdminAddSumToOrder)
//			order.Post("cancel", mw.Auth, h.AdminCancelOrder)
//			order.Post("add-receipt", mw.Auth, h.AdminAddPaymentReceipt)
//		}
//
//		cbTransfer.Post("get-file", mw.Auth, h.AdminDownloadFile)
//	}
//}
