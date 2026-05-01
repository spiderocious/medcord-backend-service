package constants

type Language string

const (
	LangEN Language = "en"
	LangES Language = "es"
	LangFR Language = "fr"
)

func LangOf(s string) Language {
	if len(s) >= 2 {
		switch Language(s[:2]) {
		case LangES:
			return LangES
		case LangFR:
			return LangFR
		}
	}
	return LangEN
}

var messages = map[MessageKey]map[Language]string{
	MsgSuccess: {
		LangEN: "Request successful",
		LangES: "Solicitud exitosa",
		LangFR: "Demande réussie",
	},
	MsgInternalServerError: {
		LangEN: "Internal server error",
		LangES: "Error interno del servidor",
		LangFR: "Erreur interne du serveur",
	},
	MsgNotFound: {
		LangEN: "Resource not found",
		LangES: "Recurso no encontrado",
		LangFR: "Ressource introuvable",
	},
	MsgInvalidPayload: {
		LangEN: "Invalid request payload",
		LangES: "Carga útil de la solicitud inválida",
		LangFR: "Charge utile de la requête invalide",
	},
	MsgValidationFailed: {
		LangEN: "Validation failed",
		LangES: "Validación fallida",
		LangFR: "Échec de la validation",
	},
	MsgUnauthorized: {
		LangEN: "Unauthorized",
		LangES: "No autorizado",
		LangFR: "Non autorisé",
	},
	MsgForbidden: {
		LangEN: "Forbidden",
		LangES: "Prohibido",
		LangFR: "Interdit",
	},
	MsgTokenRequired: {
		LangEN: "Authentication token required",
		LangES: "Se requiere token de autenticación",
		LangFR: "Jeton d'authentification requis",
	},
	MsgInvalidToken: {
		LangEN: "Invalid or expired token",
		LangES: "Token inválido o expirado",
		LangFR: "Jeton invalide ou expiré",
	},
	MsgHealthOK: {
		LangEN: "Service is healthy",
		LangES: "El servicio está saludable",
		LangFR: "Le service est sain",
	},
}

func Translate(key MessageKey, lang Language) string {
	if msgs, ok := messages[key]; ok {
		if m, ok := msgs[lang]; ok {
			return m
		}
		return msgs[LangEN]
	}
	return string(key)
}
