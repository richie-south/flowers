import * as React from 'react'
import {
  compose,
  withHandlers,
  withState,
  defaultProps,
  lifecycle
} from 'recompose'
import {withSpinner} from 'react-with-spinner'
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
export const StyledCard = styled(Card)`
  max-width: 500px !important;
  margin-top: 8px;
  margin-bottom: 8px;
`

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
        /**
         * Check has time passes
         */
        setTimeRemaining(getTimeRemaining(end))
      }, 1000)
      setIntervalId(intervalId)
    },
    componentWillUnmount() {
      const {intervalId} = this.props
      clearInterval(intervalId)
    }
  })
)

export const StatelessCountDownCard = ({
  timeRemaining: {days, hours, minutes, seconds}
}) => (
  <StyledCard>
    <StyledCardTitle
      subtitle="Next watering sessions in"
      title={`${days}d:${hours}h:${minutes}m:${seconds}s`}
    />
  </StyledCard>
)

const StyledBlockWrap = styled(Column)`
  margin-top: 16px;
  padding: 16px;
  max-width: 500px;
`

const StyledBlockTitle = styled(Title)`
  color: #444444;
  margin-bottom: 32px;
`

const StyledBlockTimeText = styled(Subhead)`
  margin-left: 4px;
  font-weight: 400;
  color: ${materialColors['grey-600']};
`
const StyledBlockTime = styled(Subhead)`color: ${materialColors['amber-500']};`

const StyledRow = styled(Row)`
  h4 {
    line-height: 28px !important;
  }
  margin-left: 16px;
`

const suffixS = (val, base) => (val === 1 ? base : `${base}s`)

export const StatelessCountDownBlock = ({
  timeRemaining: {days, hours, minutes, seconds}
}) => (
  <StyledBlockWrap>
    <StyledBlockTitle>Next watering sessions</StyledBlockTitle>
    <StyledRow>
      <StyledBlockTime>{days}</StyledBlockTime>
      <StyledBlockTimeText>{suffixS(days, 'Day')}</StyledBlockTimeText>
    </StyledRow>
    <StyledRow>
      <StyledBlockTime>{hours}</StyledBlockTime>
      <StyledBlockTimeText>{suffixS(hours, 'Hour')}</StyledBlockTimeText>
    </StyledRow>
    <StyledRow>
      <StyledBlockTime>{minutes}</StyledBlockTime>
      <StyledBlockTimeText>{suffixS(minutes, 'Minute')}</StyledBlockTimeText>
    </StyledRow>
    <StyledRow>
      <StyledBlockTime>{seconds}</StyledBlockTime>
      <StyledBlockTimeText>{suffixS(seconds, 'Second')}</StyledBlockTimeText>
    </StyledRow>
  </StyledBlockWrap>
)

const addIfNotZero = (toCheck, toAdd) => (toCheck === 0 ? '' : toAdd)

const buildTimeString = ({days, hours, minutes, seconds}) =>
  `${addIfNotZero(days, `${days}d:`)}
  ${addIfNotZero(hours, `${hours}h:`)}
  ${addIfNotZero(minutes, `${minutes}m:`)}
  ${`${seconds}s`}`

export const StatelessCountDownSpan = ({timeRemaining}) => (
  <span>{buildTimeString(timeRemaining)}</span>
)

export const CountDownSpan = enhance(StatelessCountDownSpan)
export const CountDownCard = enhance(StatelessCountDownCard)
export const CountDownBlock = enhance(StatelessCountDownBlock)
