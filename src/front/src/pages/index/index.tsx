// import React from 'react';
import { Navigate, Outlet } from 'react-router-dom';
import { HeaderMenu } from './header/index';


export const IndexLayout = () => {
    if (localStorage.getItem("AccessToken") === null) {
        return <Navigate to="/signIn" replace />
    }
    return <div>
        <HeaderMenu />

        <Outlet />
    </div>;
};