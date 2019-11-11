import React from "react";
import Styles from "./App.scss";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Checkbox from "@material-ui/core/Checkbox";
import IconButton from "@material-ui/core/IconButton";
import DeleteIcon from "@material-ui/icons/Delete";
import { makeStyles, createStyles, Theme } from "@material-ui/core/styles";

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

    deleteTodo(id: number) {
        return fetch("/api/v1/todos/" + id + "/", {
            method: "DELETE",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json"
            }
        }).then(resp => {
            if (resp.status !== 204) {
                throw new Error(resp.statusText)
            }
            return resp.json();
        }).catch(err => {
            console.error("delete todo error: ", err);
        }).then(json => {
            this.setState({
                todos: this.state.todos.filter(todo => todo.id !== id)
            });
        });
    }

    toggleTodo(id: number) {
        return fetch("/api/v1/todos/" + id + "/toggle", {
            method: "POST",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json"
            },
        }).then(resp => {
            if (resp.status !== 200) {
                throw new Error(resp.statusText)
            }
            return resp.json();
        }).catch(err => {
            console.error("toggle todo error: ", err)
        }).then(json => {
            this.setState({
                todos: this.state.todos.map(todo => {
                    return todo.id !== id ? todo : Object.assign({}, todo, { completed: !todo.completed });
                }),
            });
        });
    }

    render() {
        const todos = this.state.todos.map((t, index) => (
            <ListItem key={t.id}>
                <Checkbox checked={t.completed} onChange={() => this.toggleTodo(t.id)}/>
                <ListItemText primary={t.completed ? <s>{t.title}</s> : t.title} />
                <IconButton onClick={() => this.deleteTodo(t.id)}>
                    <DeleteIcon />
                </IconButton>
            </ListItem>
        ));

        return (
            // 本当は@material-uiのmakeStylesを使ったほうが良さそうだけど，とりあえずCSS Modulesを使用
            <div className={Styles.root}>
                <NewTodoItem postTodo={(title: string) => this.postTodo(title)}/>
                <List>
                    {todos}
                </List>
            </div>
        )
    }
}

export default App;
