import * as React from 'react'
import {compose, withHandlers, withState, defaultProps} from 'recompose'
import Layout from 'react-toolbox/lib/layout/Layout'
import Panel from 'react-toolbox/lib/layout/Panel'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import Route from 'react-router/Route'
import Switch from 'react-router/Switch'
import {FlowersOverview} from '../flowers/flowers-overview'
import {Flower} from '../flowers/flower'

const enhance = compose()

const FlowerRoutes = () => (
  <Switch>
    <Route exact path="/flowers" component={FlowersOverview} />
    <Route exact path="/flowers/:id" component={Flower} />
  </Switch>
)

export const StatelessHome = ({}) => (
  <Layout>
    <Panel style={{backgroundColor: '#FFFFFF'}}>
      <AppBar title="Flowers" />
      <FlowerRoutes />
    </Panel>
  </Layout>
)

export const Home = enhance(StatelessHome)
