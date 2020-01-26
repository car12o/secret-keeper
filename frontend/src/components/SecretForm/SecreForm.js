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

export const SecreForm = ({ hash }) => {
  const classes = useStyles()

  return (
    <Paper>
      <form className={classes.form}>
        <TextField label="Secret" color="secondary" margin="normal" />
        <TextField label="ExpireAfterViews" color="secondary" margin="normal" />
        <TextField label="ExpireAfter" color="secondary" margin="normal" />
        <Button classes={{ root: classes.marginTop }} variant="contained" color="secondary">
          create secret
        </Button>
      </form>
      {hash && <SecretProp classes={{ root: classes.marginTop }} label="Hash" value={hash} />}
    </Paper>
  )
}
