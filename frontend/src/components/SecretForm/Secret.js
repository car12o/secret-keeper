import React from "react"
import { makeStyles, Paper, TextField, Button } from "@material-ui/core"
import { SecretProp } from "./SecretProp"

const useStyles = makeStyles({
  form: {
    display: "flex",
    flexDirection: "column"
  },
  marginTop: {
    marginTop: "35px"
  }
})

export const Secret = ({ secret }) => {
  const classes = useStyles()

  return (
    <Paper>
      <form className={classes.form}>
        <TextField label="Hash" color="secondary" margin="normal" />
        <Button classes={{ root: classes.marginTop }} variant="contained" color="secondary">
          get secret
        </Button>
      </form>
      {secret && <>
        <SecretProp classes={{ root: classes.marginTop }} label="Hash" value={secret.hash} />
        <SecretProp label="SecretText" value={secret.secretText} />
        <SecretProp label="CreatedAt" value={secret.createdAt} />
        <SecretProp label="ExpiresAt" value={secret.expiresAt} />
        <SecretProp label="RemainingViews" value={secret.remainingViews} />
      </>}
    </Paper>
  )
}
