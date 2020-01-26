import React, { useState, createContext } from "react"
import { Snackbar as MatSnackbar } from "@material-ui/core"

export const SnackbarContext = createContext()

export const SnackBarProvider = ({ children }) => {
  const [open, setOpen] = useState(false)
  const [message, setMessage] = useState("")

  const handleClose = (event, reason) => {
    if (reason === "clickaway") {
      return
    }

    setOpen(false)
  }

  return (
    <SnackbarContext.Provider value={{ setOpen, setMessage }}>
      <MatSnackbar
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "left"
        }}
        open={open}
        autoHideDuration={6000}
        onClose={handleClose}
        message={message}
      />
      {children}
    </SnackbarContext.Provider>
  )
}
