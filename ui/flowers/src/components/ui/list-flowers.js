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
import ListItem from 'react-toolbox/lib/list/ListItem'
import {GET_FLOWER_LIST_URL} from '../../config/urls'

const enhance = compose(
  withState('data', 'setData', {loading: true}),
  withState('flowers', 'setFlowers', []),
  withHandlers({
    setLoading: ({setData}) => isLoading =>
      setData(data => ({
        ...data,
        loading: isLoading
      }))
  }),
  WrappedComponent => props => {
    return <WrappedComponent {...props} />
  },
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
        })
    }
  }),
  withSpinner()
)

export const StatelessListFlowers = ({flowers}) => (
  <List>
    {flowers.map(({id, name, nextWateringSession}) => (
      <ListItem
        key={id}
        caption={name}
        legend={nextWateringSession}
        leftIcon="local_florist"
      />
    ))}
  </List>
)

export const ListFlowers = enhance(StatelessListFlowers)
