/* This example requires Tailwind CSS v2.0+ */
import {Disclosure} from '@headlessui/react'
import {Bars3Icon, XMarkIcon} from '@heroicons/react/24/outline'
import {Link, LinkProps, useMatch, useResolvedPath} from "react-router-dom"
import {ToolBar} from './toolbar'
import {createContext, useEffect, useState} from "react";
import {postUserInfo} from "../../../api/sign";

const navigation = [
    {name: '博客首页', href: '/'},
    {name: '文章列表', href: '/article'},
    {name: '我的文章', href: '/myArticle'},
    {name: '个人中心', href: '/userCenter'},
    {name: '管理入口', href: '/manager'},
]

export const UserContext = createContext({id: "", username: ""})

export const HeaderMenu = () => {
    const [user, setUser] = useState({id: "", username: ""})
    useEffect(() => {
        postUserInfo().then((res) => {
            setUser(res.data)
        })
    }, [])
    return (
        <UserContext.Provider value={user}>
            <Disclosure as="nav" className="bg-gray-800">
                {({open}) => (
                    <>
                        <div className="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
                            <div className="relative flex h-16 items-center justify-between">
                                <div className="absolute inset-y-0 left-0 flex items-center sm:hidden">
                                    {/* Mobile menu button*/}
                                    <Disclosure.Button
                                        className="inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white">
                                        <span className="sr-only">主菜单</span>
                                        {open ? (
                                            <XMarkIcon className="block h-6 w-6" aria-hidden="true"/>
                                        ) : (
                                            <Bars3Icon className="block h-6 w-6" aria-hidden="true"/>
                                        )}
                                    </Disclosure.Button>
                                </div>
                                <div
                                    className="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
                                    <div className="flex flex-shrink-0 items-center">
                                        <img
                                            className="block h-8 w-auto lg:hidden"
                                            src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
                                            alt="Your Company"
                                        />
                                        <img
                                            className="hidden h-8 w-auto lg:block"
                                            src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
                                            alt="Your Company"
                                        />
                                    </div>
                                    <div className="hidden sm:ml-6 sm:block">
                                        <div className="flex space-x-4">
                                            {navigation.map((item) => (
                                                <CustomLink
                                                    key={item.name}
                                                    to={item.href}
                                                >
                                                    {item.name}
                                                </CustomLink>
                                            ))}
                                        </div>
                                    </div>
                                </div>
                                <ToolBar/>
                            </div>
                        </div>
                        <Disclosure.Panel className="sm:hidden">
                            <div className="space-y-1 px-2 pt-2 pb-3">
                                {navigation.map((item) => (
                                    <Disclosure.Button
                                        key={item.name}
                                        as={CustomLink}
                                        to={item.href}
                                    >
                                        {item.name}
                                    </Disclosure.Button>
                                ))}
                            </div>
                        </Disclosure.Panel>
                    </>
                )}
            </Disclosure>
        </UserContext.Provider>


    )
}

export function classNames(...classes: string[]) {
    return classes.filter(Boolean).join(' ')
}

const CustomLink = ({children, to, ...props}: LinkProps) => {
    let resolved = useResolvedPath(to)
    let match = useMatch({path: resolved.pathname, end: true})
    return (
        <div>
            <Link
                className={classNames(
                    match ? 'bg-gray-900 text-white' : 'text-gray-300 hover:bg-gray-700 hover:text-white',
                    'block px-3 py-2 rounded-md text-base font-medium'
                )}
                to={to}
                aria-current={match ? 'page' : undefined}
                {...props}
            >
                {children}
            </Link>
        </div>
    )
}
