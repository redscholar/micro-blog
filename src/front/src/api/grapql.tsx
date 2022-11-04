import {graphql,buildSchema} from 'graphql'

var schema = buildSchema(`type Query {
    hello: String
}`)

var rootValue = {
    hello: ()=> {
        return "hello world"
    }
}