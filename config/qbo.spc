connection "qbo" {
    plugin = "qbo"

    # The base URL to call for access to the QBO API.
    baseURL = ""

    # Client ID issued by the QBO developer portal.
    clientId = ""

    # Client Secret issued by the QBO developer portal.
    clientSecret = ""

    # Realm ID issued by the QBO developer portal. This is equivalent
    # to the Company ID, and the terms are used interchangably.
    realmId = ""

    # The initial refresh token from the QBO developer portal. This will
    # need to be refreshed regularly, usually every 101 days
    refreshToken = ""
}
