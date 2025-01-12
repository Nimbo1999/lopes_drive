import { ButtonHTMLAttributes } from "react";
import clsx from "clsx";

import styles from "./button.module.scss";

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "default" | "small";
}

export function Button({
  children,
  className,
  variant = "default",
  ...props
}: Readonly<ButtonProps>) {
  return (
    <button
      {...props}
      className={clsx(className, styles.button, {
        [styles["button--small"]]: variant === "small",
      })}
    >
      {children}
    </button>
  );
}
