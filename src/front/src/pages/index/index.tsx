// import React from 'react';
import {Navigate, Outlet} from 'react-router-dom';
import {MessageLayout} from '../component/message';
import {HeaderMenu} from './header/index';
import {postUserInfo} from "../../api/sign";
import {createContext, useState} from "react";

export const UserContext = createContext({id: "", username: ""})

export const IndexLayout = () => {
    const [user, setUser] = useState({id: "", username: ""})
    if (localStorage.getItem("AccessToken") === null) {
        return <Navigate to="/signIn" replace/>
    } else {
        postUserInfo().then((res: { id: string, username: string }) => {
            setUser(res)
        })
    }

    return (
        <UserContext.Provider value={user}>
            <div>
                <HeaderMenu/>


                <MessageLayout/>
                <Outlet/>
            </div>
        </UserContext.Provider>
    )
};