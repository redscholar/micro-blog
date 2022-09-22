import { Navigate } from 'react-router-dom';

export const SignOut = () => {
    localStorage.removeItem("AccessToken")
    return <Navigate to="/signIn" replace />
}