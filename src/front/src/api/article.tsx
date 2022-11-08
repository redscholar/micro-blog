import {postGraphql} from "./request";

interface createArticleRequest {
    title: String
    content: String
    image: String
}

export function createArticle(param: createArticleRequest) {
    const query = {
        operationName: "createArticle",
        query: `mutation createArticle($request: createArticleRequest!) {
    createArticle(request: $request) {
        code
        msg
    }
}`,
        variables: {
            request: param
        }
    }
    return postGraphql("/query", query)
}


interface listArticleRequest {
    keyword: String
    lastId: String
    pagination: Pagination
}

interface Pagination {
    page: number
    limit: number
}

export function listArticle(param: listArticleRequest) {
    const query = {
        operationName: "listArticle",
        query: `query listArticle($request: listArticleRequest!) {
    listArticle(request: $request) {
        code
        msg
        data{
            total
            articles{
                id
                title
                content
                createdAt
            }
        }
    }
}`,
        variables: {
            request: param
        }
    }
    return postGraphql("/query", query)
}