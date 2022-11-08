import {ChevronLeftIcon, ChevronRightIcon} from '@heroicons/react/20/solid'
import {Link, Outlet} from "react-router-dom";
import {useEffect, useState} from "react";
import {listArticle} from "../../../api/article";


export const MyArticle = () => {
    return (
        < Outlet/>
    )

}

export const SearchMyArticle = () => {
    const [articles, setArticles] = useState([])
    const [page, setPage] = useState([{
        total: 0,
        page: 1,
        limit: 5
    }])
    const [searchReq, setSearchReq] = useState({
        keyword: "bbb",
        lastId: "",
        pagination: {
            page: 1,
            limit: 5
        },
    })
    const search = () => {
        listArticle(searchReq).then((res: any) => {
            setArticles(res.data.articles)
        })
    }
    useEffect(() => {
        search()
    }, [])

    return (
        <>
            <div className="container mx-auto flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
                <div className="relative rounded-lg bg-white shadow-md w-full max-w-2xl">
                    <input type="text" name="Search" placeholder="Search..."
                           className="w-full text-xl py-4 pl-10 focus:outline-none pr-10"
                           onChange={(e) => setSearchReq(() => {
                               searchReq.keyword = e.target.value
                               return searchReq
                           })}/>
                    <span className="absolute top-4 right-4 h-6 w-6 fill-slate-400">
                        <button className="h-6 w-6 focus:outline-none focus:ring" onClick={search}>
                            <svg className="h-6 w-6"
                                 xmlns="http://www.w3.org/2000/svg">
                                    <path
                                        d="M20.47 21.53a.75.75 0 1 0 1.06-1.06l-1.06 1.06Zm-9.97-4.28a6.75 6.75 0 0 1-6.75-6.75h-1.5a8.25 8.25 0 0 0 8.25 8.25v-1.5ZM3.75 10.5a6.75 6.75 0 0 1 6.75-6.75v-1.5a8.25 8.25 0 0 0-8.25 8.25h1.5Zm6.75-6.75a6.75 6.75 0 0 1 6.75 6.75h1.5a8.25 8.25 0 0 0-8.25-8.25v1.5Zm11.03 16.72-5.196-5.197-1.061 1.06 5.197 5.197 1.06-1.06Zm-4.28-9.97c0 1.864-.755 3.55-1.977 4.773l1.06 1.06A8.226 8.226 0 0 0 18.75 10.5h-1.5Zm-1.977 4.773A6.727 6.727 0 0 1 10.5 17.25v1.5a8.226 8.226 0 0 0 5.834-2.416l-1.061-1.061Z"></path>
                            </svg>
                        </button>
                    </span>
                </div>
                <Link to="/myArticle/add"
                      className="text-sm relative pl-10 pt-6 text-sky-400 hover:text-blue-500">新增文章&gt;&gt;</Link>
            </div>
            <div className="container p-2 mx-auto sm:p-4 dark:text-gray-100">
                <div className="grid justify-center grid-cols-1 gap-6">
                    {articles.map((item: any) => {
                        return (
                            <a key={item.id} rel="noopener noreferrer" href="#"
                               className="hover:bg-slate-50 group hover:no-underline focus:no-underline dark:bg-gray-900">
                                <div className="p-6 space-y-2">
                                    <h3 className="text-2xl font-semibold group-focus:underline">{item.title}</h3>
                                    <span className="text-xs dark:text-gray-400">{item.createdAt}</span>
                                    <p className="truncate">{item.content}</p>
                                </div>
                            </a>
                        )
                    })}
                </div>


                <div className="flex items-center justify-between border-t border-gray-200 bg-white px-4 py-3 sm:px-6">
                    <div className="flex flex-1 justify-between sm:hidden">
                        <a
                            href="#"
                            className="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
                        >
                            Previous
                        </a>
                        <a
                            href="#"
                            className="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
                        >
                            Next
                        </a>
                    </div>
                    <div className="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
                        <div>
                            <p className="text-sm text-gray-700">
                                Showing <span className="font-medium">1</span> to <span
                                className="font-medium">10</span> of{' '}
                                <span className="font-medium">97</span> results
                            </p>
                        </div>
                        <div>
                            <nav className="isolate inline-flex -space-x-px rounded-md shadow-sm"
                                 aria-label="Pagination">
                                <a
                                    href="#"
                                    className="relative inline-flex items-center rounded-l-md border border-gray-300 bg-white px-2 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20"
                                >
                                    <span className="sr-only">Previous</span>
                                    <ChevronLeftIcon className="h-5 w-5" aria-hidden="true"/>
                                </a>
                                {/* Current: "z-10 bg-indigo-50 border-indigo-500 text-indigo-600", Default: "bg-white border-gray-300 text-gray-500 hover:bg-gray-50" */}
                                <a
                                    href="#"
                                    aria-current="page"
                                    className="relative z-10 inline-flex items-center border border-indigo-500 bg-indigo-50 px-4 py-2 text-sm font-medium text-indigo-600 focus:z-20"
                                >
                                    1
                                </a>
                                <a
                                    href="#"
                                    className="relative inline-flex items-center border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20"
                                >
                                    2
                                </a>
                                <a
                                    href="#"
                                    className="relative hidden items-center border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20 md:inline-flex"
                                >
                                    3
                                </a>
                                <span
                                    className="relative inline-flex items-center border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700">
              ...
            </span>
                                <a
                                    href="#"
                                    className="relative hidden items-center border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20 md:inline-flex"
                                >
                                    8
                                </a>
                                <a
                                    href="#"
                                    className="relative inline-flex items-center border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20"
                                >
                                    9
                                </a>
                                <a
                                    href="#"
                                    className="relative inline-flex items-center border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20"
                                >
                                    10
                                </a>
                                <a
                                    href="#"
                                    className="relative inline-flex items-center rounded-r-md border border-gray-300 bg-white px-2 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20"
                                >
                                    <span className="sr-only">Next</span>
                                    <ChevronRightIcon className="h-5 w-5" aria-hidden="true"/>
                                </a>
                            </nav>
                        </div>
                    </div>
                </div>
            </div>
        </>

    )
}

