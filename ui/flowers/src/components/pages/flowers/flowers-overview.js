import * as React from 'react'
import {
  compose,
  withHandlers,
  withState,
  defaultProps,
  lifecycle
} from 'recompose'
import {withSpinner} from 'react-with-spinner'
import List from 'react-toolbox/lib/list/List'
import Redirect from 'react-router/Redirect'
import ListItem from 'react-toolbox/lib/list/ListItem'
import {GET_FLOWER_LIST_URL} from '../../../config/urls'
import {injectIntl} from 'react-intl'
import {ListFlowers} from '../../ui/list-flowers'

const enhance = compose(
  withState('data', 'setData', {loading: true}),
  withState('flowers', 'setFlowers', []),
  withState('redirect', 'setRedirect', ''),
  withHandlers({
    setLoading: ({setData}) => isLoading =>
      setData(data => ({
        ...data,
        loading: isLoading
      })),
    viewFlower: ({setRedirect}) => id => setRedirect(id)
  }),
  lifecycle({
    componentDidMount(props) {
      fetch(GET_FLOWER_LIST_URL)
        .then(response => response.json())
        .then(flowers => {
          this.props.setFlowers(flowers)
          this.props.setLoading(false)
        })
        .catch(error => {
          // handle error
          console.log(error)
        })
    }
  }),
  withSpinner(),
  injectIntl
)

export const StatelessFlowersOverview = ({
  flowers,
  intl,
  redirect,
  viewFlower,
}) =>
  redirect ? (
    <Redirect to={redirect} />
  ) : (
    <div>
      <h1>Flowers list</h1>
      <ListFlowers viewFlower={viewFlower} flowers={flowers} />
    </div>
  )

export const FlowersOverview = enhance(StatelessFlowersOverview)
