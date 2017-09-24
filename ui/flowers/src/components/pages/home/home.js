import * as React from 'react'
import {compose, withHandlers, withState, defaultProps} from 'recompose'
import Layout from 'react-toolbox/lib/layout/Layout'
import Panel from 'react-toolbox/lib/layout/Panel'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import {ListFlowers} from '../../ui/list-flowers'

const enhance = compose(
)

export const StatelessHome = ({}) => (
  <Layout>
    <Panel>
      <AppBar title="Flowers" />
      <h1>HOME</h1>
      <ListFlowers />
    </Panel>
  </Layout>
)

export const Home = enhance(StatelessHome)
