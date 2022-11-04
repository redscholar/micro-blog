// import React from 'react';
import {BrowserRouter, Route, Routes} from 'react-router-dom';
// import './App.css';
import {SignIn} from './pages/sign/SignIn'
import {SignOut} from './pages/sign/SignOut'
import {ErrorPage404} from './pages/error/404'
import {Article} from './pages/index/article';
import {Blog} from './pages/index/blog';
import {MyArticle, SearchMyArticle} from './pages/index/myArticle';
import {UserCenter} from './pages/index/user';
import {Manager} from './pages/index/manager';
import {IndexLayout} from './pages/index';
import SignUp from './pages/sign/signUp';
import {AddArticle} from "./pages/index/myArticle/addArticle";

export default function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/signIn" element={<SignIn/>}/>
                <Route path="/signUp" element={<SignUp/>}/>
                <Route path="/signOut" element={<SignOut/>}/>
                <Route path="/" element={<IndexLayout/>}>
                    <Route index element={<Blog/>}/>
                    <Route path="index" element={<Blog/>}/>
                    <Route path="/article" element={<Article/>}/>
                    <Route path="/myArticle" element={<MyArticle/>}>
                        <Route index element={<SearchMyArticle/>} />
                        <Route path="add" element={<AddArticle/>}/>
                    </Route>
                    <Route path="/userCenter" element={<UserCenter/>}/>
                    <Route path="/manager" element={<Manager/>}/>
                </Route>
                <Route path="*" element={<ErrorPage404/>}/>
            </Routes>
        </BrowserRouter>
    );
}

