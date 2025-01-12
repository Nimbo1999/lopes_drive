import { ButtonHTMLAttributes } from "react";
import clsx from "clsx";

import styles from "./button.module.scss";

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {}

export function Button({
  children,
  className,
  ...props
}: Readonly<ButtonProps>) {
  return (
    <button {...props} className={clsx(className, styles.button)}>
      {children}
    </button>
  );
}
