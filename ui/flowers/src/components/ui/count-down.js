import * as React from 'react'
import {
  compose,
  withHandlers,
  withState,
  defaultProps,
  lifecycle
} from 'recompose'
import {withSpinner} from 'react-with-spinner'
import {injectIntl} from 'react-intl'
import Avatar from 'react-toolbox/lib/avatar/Avatar'
import IconButton from 'react-toolbox/lib/button/IconButton'
import Button from 'react-toolbox/lib/button/Button'
import WaterPumpIcon from 'mdi-react/WaterPumpIcon'
import {Column, Row} from 'styled-material/dist/src/layout'
import styled from 'styled-components'
import {Caption, Subhead, Title} from 'styled-material/dist/src/typography'
import {materialColors} from 'styled-material/dist/src/colors'
import Card from 'react-toolbox/lib/card/Card'
import CardMedia from 'react-toolbox/lib/card/CardMedia'
import CardTitle from 'react-toolbox/lib/card/CardTitle'
import CardText from 'react-toolbox/lib/card/CardText'
import CardActions from 'react-toolbox/lib/card/CardActions'

export const StyledCardTitle = styled(CardTitle)`
  justify-content: center;
  h5 {
    color: ${materialColors['amber-500']};
  }
`
export const StyledCard = styled(Card)`max-width: 500px !important;`

const getTimeRemaining = endtime => {
  const total = Date.parse(endtime) - Date.parse(new Date())
  const seconds = Math.floor((total / 1000) % 60)
  const minutes = Math.floor((total / 1000 / 60) % 60)
  const hours = Math.floor((total / (1000 * 60 * 60)) % 24)
  const days = Math.floor(total / (1000 * 60 * 60 * 24))
  return {
    total,
    days,
    hours,
    minutes,
    seconds
  }
}

const enhance = compose(
  withState('intervalId', 'setIntervalId'),
  withState('timeRemaining', 'setTimeRemaining', {
    total: 0,
    days: 0,
    hours: 0,
    minutes: 0,
    seconds: 0
  }),
  lifecycle({
    componentDidMount() {
      const {setIntervalId, setTimeRemaining, timeRemaining} = this.props
      const intervalId = setInterval(() => {
        const {end} = this.props
        setTimeRemaining(getTimeRemaining(end))
      }, 1000)
      setIntervalId(intervalId)
    },
    componentWillUnmount() {
      const {intervalId} = this.props
      clearInterval(intervalId)
    }
  }),
  injectIntl
)

export const StatelessCountDownCard = ({
  timeRemaining: {days, hours, minutes, seconds}
}) => (
  <StyledCard>
    <StyledCardTitle
      subtitle='Next watering sessions in:'
      title={`${days}d:${hours}h:${minutes}m:${seconds}s`}
    />
  </StyledCard>
)

export const StatelessCountDownSpan = ({
  timeRemaining: {days, hours, minutes, seconds}
}) => <span>{`${days}d:${hours}h:${minutes}m:${seconds}s`}</span>

export const CountDownSpan = enhance(StatelessCountDownSpan)
export const CountDownCard = enhance(StatelessCountDownCard)
