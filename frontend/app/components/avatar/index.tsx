import clsx from "clsx";
import { HTMLAttributes, ImgHTMLAttributes } from "react";

import styles from "./avatar.module.scss";

export interface AvatarProps
  extends Omit<HTMLAttributes<HTMLDivElement>, "children"> {
  src: ImgHTMLAttributes<HTMLImageElement>["src"];
  alt: ImgHTMLAttributes<HTMLImageElement>["alt"];
}

export function Avatar({
  src,
  alt,
  className,
  ...props
}: Readonly<AvatarProps>) {
  return (
    <div {...props} className={clsx(className, styles.avatar)}>
      <img src={src} alt={alt} className={styles.avatar__image} />
    </div>
  );
}
