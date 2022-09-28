// import React from 'react';
import {Navigate, Outlet} from 'react-router-dom';
import {MessageLayout} from '../component/message';
import {HeaderMenu} from './header';

export const IndexLayout = () => {
    if (localStorage.getItem("AccessToken") === null) {
        return <Navigate to="/signIn" replace/>
    }

    return (
        <div>
            <HeaderMenu/>

            <MessageLayout/>
            <Outlet/>
        </div>
    )
};