import React from 'react'

const UserAuthContext = React.createContext({id: '', username: '', name: ''})
export const UserAuthProvider = UserAuthContext.Provider
export const UserAuthConsumer = UserAuthContext.Consumer

export default UserAuthContext