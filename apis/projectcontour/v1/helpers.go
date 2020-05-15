package v1

// GetPrefixReplacements returns replacement prefixes from the path
// rewrite policy (if any).
func (r *Route) GetPrefixReplacements() []ReplacePrefix {
	if r.PathRewritePolicy != nil {
		return r.PathRewritePolicy.ReplacePrefix
	}
	return nil
}

// DisableAuthentication returns whether Contour should disable
// request authorization on this route.
func (r *Route) DisableAuthorization() bool {
	if r.AuthPolicy != nil {
		return r.AuthPolicy.Disabled
	}

	return false
}

// AuthorizationContext merges the parent context entries with the context from this Route.
func (r *Route) AuthorizationContext(parent map[string]string) map[string]string {
	values := make(map[string]string, len(parent))

	for k, v := range parent {
		values[k] = v
	}

	if r.AuthPolicy != nil {
		for k, v := range r.AuthPolicy.Context {
			values[k] = v
		}
	}

	if len(values) == 0 {
		return nil
	}

	return values
}

// AuthorizationConfigured returns whether authorization  is
// configured on this virtual host.
func (v *VirtualHost) AuthorizationConfigured() bool {
	return v.TLS != nil && v.Authorization != nil
}

// DisableAuthorization returns true if this virtual host disables
// authorization. If an authorization server is present, the the default
// policy is to not disable.
func (v *VirtualHost) DisableAuthorization() bool {
	// No authorization, so it is disabled.
	if v.AuthorizationConfigured() {
		// No policy specified, default is to not disable.
		if v.Authorization.AuthPolicy == nil {
			return false
		}

		return v.Authorization.AuthPolicy.Disabled
	}

	return false
}

// AuthorizationContext returns the authorization policy context (if present).
func (v *VirtualHost) AuthorizationContext() map[string]string {
	if v.AuthorizationConfigured() {
		if v.Authorization.AuthPolicy != nil {
			return v.Authorization.AuthPolicy.Context
		}
	}

	return nil
}
