import React from 'react';
export const ErrorPage404 = () => {
    return <section className="flex items-center h-full p-16 dark:bg-gray-900 dark:text-gray-100">
        <div className="container flex flex-col items-center justify-center px-5 mx-auto my-8">
            <div className="max-w-md text-center">
                <h2 className="mb-8 font-extrabold text-9xl dark:text-gray-600">
                    <span className="sr-only">Error</span>404
                </h2>
                {/* <p className="text-2xl font-semibold md:text-3xl">页面缺失</p> */}
                <p className="mt-4 mb-8 dark:text-gray-400">页面缺失</p>
                <a rel="noopener noreferrer" onClick={backToHomepage} className="px-8 py-3 font-semibold rounded dark:bg-violet-400 dark:text-gray-900">返回主页</a>
            </div>
        </div>
    </section>
}

const backToHomepage = () => {
    window.location.href = "/"
}