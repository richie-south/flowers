import React, {Component} from 'react'
import {BrowserRouter} from 'react-router-dom'
import ThemeProvider from 'react-toolbox/lib/ThemeProvider'
import {TopRouter} from './components/top-router'
import {client, store} from './lib/store'
import theme from './styles/theme'
import './styles/theme.css'
import './App.css'

class App extends Component {
  render() {
    return (
      <ThemeProvider theme={theme}>
        <BrowserRouter>
          <TopRouter />
        </BrowserRouter>
      </ThemeProvider>
    )
  }
}

export default App
