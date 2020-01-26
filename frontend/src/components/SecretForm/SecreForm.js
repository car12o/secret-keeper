import React, { useState, useContext } from "react"
import { makeStyles, Paper, TextField, Button } from "@material-ui/core"
import { SnackbarContext } from "../Snackbar"
import { SecretProp } from "./SecretProp"
import { post } from "../../api"


const useStyles = makeStyles({
  form: {
    display: "flex",
    flexDirection: "column"
  },
  marginTop: {
    marginTop: "35px"
  }
})

export const SecreForm = () => {
  const [data, setData] = useState({
    secret: "",
    expireAfterViews: "",
    expireAfter: ""
  })
  const [hash, setHash] = useState("")
  const classes = useStyles()
  const { setOpen, setMessage } = useContext(SnackbarContext)

  const postSecret = async () => {
    const { status, data: resp } = await post(data)
    if (status >= 300) {
      setMessage(JSON.stringify(resp))
      setOpen(true)
      return
    }

    setHash(resp.hash)
  }

  const handleChange = (key) => ({ target: { value } }) => {
    const int = parseInt(value, 10)
    if (int) {
      setData({ ...data, [key]: int })
      return
    }

    setData({ ...data, [key]: value })
  }

  return (
    <Paper>
      <form className={classes.form}>
        <TextField label="Secret" color="secondary" margin="normal"
          value={data.secret} onChange={handleChange("secret")}
        />
        <TextField type="number" label="ExpireAfterViews" color="secondary" margin="normal"
          value={data.expireAfterViews} onChange={handleChange("expireAfterViews")}
        />
        <TextField type="number" label="ExpireAfter" color="secondary" margin="normal"
          value={data.expireAfter} onChange={handleChange("expireAfter")}
        />
        <Button
          classes={{ root: classes.marginTop }}
          variant="contained"
          color="secondary"
          onClick={postSecret}
        >
          create secret
        </Button>
      </form>
      {hash && <SecretProp classes={{ root: classes.marginTop }} label="Hash" value={hash} />}
    </Paper>
  )
}
