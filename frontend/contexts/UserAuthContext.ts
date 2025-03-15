import {createContext} from 'react'

const UserAuthContext = createContext({id: '', username: '', name: ''})
export const UserAuthProvider = UserAuthContext.Provider
export const UserAuthConsumer = UserAuthContext.Consumer

export default UserAuthContext