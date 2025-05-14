
const regexes = {
  password: {
    strong: /^(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[^\w\d\s:])([^\s]){8,}$/,
    medium: /^(?=.*\d)(?=.*[A-Z])(?=.*[a-z])([^\s]){6,}$/,
    weak: /^(?=.*[a-z])([^\s]){8,}$/,
  }
} as const

type Regexes = typeof regexes


export { regexes, type Regexes }
