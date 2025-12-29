export type T_Session = {
	id: string;
	name?: string;
	email: string;
	picture: string;
};

class SessionStore {
	session = $state<T_Session | null>(null);

	private parseJwt(token: string) {
		return JSON.parse(atob(token.split('.')[1]));
	}

	init() {
		if (this.session !== null) return;

		const token = localStorage.getItem(import.meta.env.APP_TOKEN_ISSUER);
		if (token) {
			try {
				const payload = this.parseJwt(token);
				this.session = payload.user as T_Session;
				console.log('Session initialized:', this.session);
			} catch (error) {
				console.error('Failed to parse JWT:', error);
				this.clear();
			}
		}
	}

	get() {
		return this.session;
	}

	set(newSession: T_Session | null) {
		this.session = newSession;
	}

	clear() {
		this.session = null;
		localStorage.removeItem(import.meta.env.APP_TOKEN_ISSUER);
	}

	isAuthenticated() {
		return this.session !== null;
	}
}

export const sessionStore = new SessionStore();

export function setToken() {
	return { authorization: 'bearer ' + localStorage.getItem(import.meta.env.APP_TOKEN_ISSUER) };
}
