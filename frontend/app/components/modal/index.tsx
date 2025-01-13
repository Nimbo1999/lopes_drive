import { createPortal } from "react-dom";
import { HTMLAttributes, useRef } from "react";
import clsx from "clsx";
import { useClickOutside } from "~/hooks";

import { Button } from "../button";
import styles from "./modal.module.scss";

export interface ModalProps extends HTMLAttributes<HTMLDivElement> {
  isOpen?: boolean;
  onClose?: VoidFunction;
}

export function Modal({
  children,
  className,
  isOpen = false,
  onClose,
  ...props
}: Readonly<ModalProps>) {
  const dialogRef = useRef<HTMLDivElement>(null);
  useClickOutside(dialogRef, onClose);

  if (typeof document === "undefined") return null;

  const modal = (
    <div
      className={clsx(styles.modal, className, { [styles.open]: isOpen })}
      {...props}
    >
      <div ref={dialogRef} role="dialog" className={styles.modal__dialog}>
        <div className={styles.modal__content}>{children}</div>

        {typeof onClose === "function" ? (
          <footer className={styles.modal__footer}>
            <Button fullWidth type="button" onClick={onClose} variant="text">
              Close
            </Button>
          </footer>
        ) : null}
      </div>
    </div>
  );
  return createPortal(modal, document.getElementById("react-portal")!);
}
