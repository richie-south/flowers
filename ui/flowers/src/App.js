import React, {Component} from 'react'
import {IntlProvider} from 'react-intl'
import {BrowserRouter} from 'react-router-dom'
import ThemeProvider from 'react-toolbox/lib/ThemeProvider'
import {TopRouter} from './components/top-router'
// import {client, store} from './lib/store'
import theme from './styles/theme'
import './styles/theme.css'
import './App.css'
import 'isomorphic-fetch'
import * as sv from './translations/sv.json'

class App extends Component {
  render() {
    return (
      <ThemeProvider theme={theme}>
        <IntlProvider locale="sv" messages={sv}>
          <BrowserRouter>
            <TopRouter />
          </BrowserRouter>
        </IntlProvider>
      </ThemeProvider>
    )
  }
}

export default App
