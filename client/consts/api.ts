export const API = {
  restcountries: {
    base: 'https://restcountries.com/v3.1',
    alpha: 'https://restcountries.com/v3.1/alpha'
  },
  bitacoraForum: {
    base: 'http://127.21.0.11:8080/api/v1',
    auth: {
      register: 'http://127.21.0.11:8080/api/v1/auth/register',
      login: 'http://127.21.0.11:8080/api/v1/auth/login',
    },
    posts: {
      all: 'http://127.21.0.11:8080/api/v1/post/all',
      countCountry: 'http://127.21.0.11:8080/api/v1/post/count/',
    }
  },
}
