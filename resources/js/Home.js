import React from "react";
import { Switch, Route, Link } from "react-router-dom";
import ViewList from "./components/ViewList";

export default function Home() {
  return (
    <div>
      <Switch>
        <Route exact path="/" render={() => <ViewList />} />
      </Switch>
    </div>
  );
}
