import React from "react";
import Styles from "./App.scss";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Checkbox from "@material-ui/core/Checkbox";

import TodoModel from "../models/todoModel";
import NewTodoItem from "./NewTodoItem"

interface AppProps {
}

interface AppState {
    todos: TodoModel[];
}

class App extends React.Component<AppProps, AppState> {
    constructor(props: AppProps) {
        super(props);

        this.state = {
            todos: [],
        };

        this.load();
    }

    load() {
        return fetch("/api/v1/todos", {
            method: "GET",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json"
            }
        }).then(resp => {
            if (resp.status !== 200) {
                throw new Error(resp.statusText)
            }
            return resp.json();
        }).catch(err => {
            console.error("get todo error: ", err);
        }).then(json => {
            this.setState({todos: json});
        });
    }

    postTodo(title: string) {
        if (title === "") {
            return;
        }

        const newTodo = {
            title: title,
            completed: false,
        };

        fetch("/api/v1/todos", {
            method: "POST",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json"
            },
            body: JSON.stringify(newTodo),
        }).then(resp => {
            if (resp.status !== 201) {
                throw new Error(resp.statusText);
            }
            return resp.json();
        }).catch(err => {
            console.error("post todo error: ", err);
        }).then(json => {
            if (json === null) {
                return;
            }

            this.setState({
                todos: this.state.todos.concat(json),
            })
        });
    }

    render() {
        const todos = this.state.todos.map((t, index) => (
            <ListItem key={t.id}>
                <Checkbox />
                <ListItemText primary={t.title} />
            </ListItem>
        ));

        return (
            <div>
                <NewTodoItem postTodo={(title: string) => this.postTodo(title)}/>
                <List>
                    {todos}
                </List>
            </div>
        )
    }
}

export default App;
