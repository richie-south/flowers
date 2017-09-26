import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import App from './App'
import registerServiceWorker from './registerServiceWorker'
import {addLocaleData} from 'react-intl'
import sv from 'react-intl/locale-data/sv'

addLocaleData(sv)

ReactDOM.render(<App />, document.getElementById('root'))
registerServiceWorker()
