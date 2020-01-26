import React from "react"
import { makeStyles, Typography } from "@material-ui/core"

const useStyles = makeStyles({
  root: {
    display: "flex",
    flexDirection: "row"
  },
  typography: {
    marginRight: "8px"
  }
})

export const SecretProp = ({ classes = {}, label, value }) => {
  const _classes = useStyles()

  return (
    <div className={`${_classes.root} ${classes.root}`}>
      <Typography classes={{ root: _classes.typography }} variant="body1" color="secondary">
        {label}
      </Typography>
      <Typography variant="body1">
        {value}
      </Typography>
    </div>
  )
}
