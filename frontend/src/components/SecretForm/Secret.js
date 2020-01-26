import React, { useState, useContext } from "react"
import { makeStyles, Paper, TextField, Button } from "@material-ui/core"
import { SnackbarContext } from "../Snackbar"
import { SecretProp } from "./SecretProp"
import { get } from "../../api"

const useStyles = makeStyles({
  form: {
    display: "flex",
    flexDirection: "column"
  },
  marginTop: {
    marginTop: "35px"
  }
})

export const Secret = () => {
  const [hash, setHash] = useState("")
  const [data, setData] = useState({
    hash: "",
    secretText: "",
    createdAt: "",
    expiresAt: "",
    remainingViews: ""
  })
  const classes = useStyles()
  const { setOpen, setMessage } = useContext(SnackbarContext)

  const getSecret = async () => {
    const { status, data: resp } = await get(hash)
    if (status >= 300) {
      setMessage(`${status}-${JSON.stringify(resp)}`)
      setOpen(true)
      return
    }

    setData(resp)
  }

  return (
    <Paper>
      <form className={classes.form}>
        <TextField label="Hash" color="secondary" margin="normal"
          value={hash} onChange={({ target: { value } }) => setHash(value)}
        />
        <Button
          classes={{ root: classes.marginTop }}
          variant="contained"
          color="secondary"
          onClick={getSecret}
        >
          get secret
        </Button>
      </form>
      {data.hash && <>
        <SecretProp classes={{ root: classes.marginTop }} label="Hash" value={data.hash} />
        <SecretProp label="SecretText" value={data.secretText} />
        <SecretProp label="CreatedAt" value={data.createdAt} />
        <SecretProp label="ExpiresAt" value={data.expiresAt} />
        <SecretProp label="RemainingViews" value={data.remainingViews} />
      </>}
    </Paper>
  )
}
