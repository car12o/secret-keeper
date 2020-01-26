import React from "react"
import { makeStyles, Typography } from "@material-ui/core"
import { SecreForm, Secret } from "./components/SecretForm"

const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
    justifyContent: "center",
    margin: theme.spacing(1, 5)
  },
  container: {
    "& > *": {
      width: "750px",
      padding: theme.spacing(4),
      marginBottom: "35px"
    }
  },
  typography: {
    padding: 0,
    marginBottom: "10px"
  }
}))

export const App = () => {
  const classes = useStyles()

  return (
    <div className={classes.root}>
      <div className={classes.container}>
        <Typography classes={{ root: classes.typography }} variant="h3" color="secondary">
          Secret
        </Typography>
        <SecreForm />
        <Secret />
      </div>
    </div>
  )
}
