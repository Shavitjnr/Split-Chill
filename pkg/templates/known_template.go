package templates

type KnownTemplate string


const (
	TEMPLATE_VERIFY_EMAIL                   KnownTemplate = "email/verify_email"
	TEMPLATE_PASSWORD_RESET                 KnownTemplate = "email/password_reset"
	SYSTEM_PROMPT_RECEIPT_IMAGE_RECOGNITION KnownTemplate = "prompt/receipt_image_recognition"
)
