import React from "react";
import Styles from "./App.scss";
import TodoModel from "../models/todoModel"

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

    render() {
        const todos = this.state.todos.map((t, index) => (
            <div key={t.id}>{t.title}</div>
        ))
        return (
            <div>
                {todos}
            </div>
        )
    }
}

export default App;
