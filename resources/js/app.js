import React from "react";
import { render } from "react-dom";
import { BrowserRouter } from "react-router-dom/cjs/react-router-dom.min";
import Home from "./Home";
import { ChakraProvider } from "@chakra-ui/react";
render(
  <BrowserRouter>
    <ChakraProvider>
      <Home />
    </ChakraProvider>
  </BrowserRouter>,
  document.getElementById("app")
);
