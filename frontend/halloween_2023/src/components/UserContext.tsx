import React, {useContext, createContext} from 'react'

export interface User  {
    username: string;
    emailaddress:string;
    pumpkinsAchieved: boolean[]
}

export interface UserContextProps {
    readonly user: User | null;
    readonly setUser: (user: (prevState: User) => User) => void
}
//ype '(prevState: User) => User' is not assignable to type '(prevState: User | null) => User | null'.


export const UserContext = React.createContext<UserContextProps>({
    user: null,
    setUser: (User) => null
})